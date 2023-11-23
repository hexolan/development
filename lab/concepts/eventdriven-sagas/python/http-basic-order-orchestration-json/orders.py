import asyncio
from enum import Enum
from typing import Optional

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel


CUSTOMER_COMMAND_CHANNEL = "http://localhost:5000/customer_commands"


###


class OrderStatus(str, Enum):
    pending = "pending"
    rejected = "rejected"
    approved = "approved"


class Order(BaseModel):
    id: int
    status: OrderStatus
    customer_id: int


###


class PlaceOrderRequest(BaseModel):
    customer_id: int


###


class CustomerCommandType(str, Enum):
    reserve_credit = "reserve_credit"


class CustomerCommandData(BaseModel):
    order_id: int
    customer_id: int


class CustomerCommand(BaseModel):
    type: CustomerCommandType
    data: CustomerCommandData


###


class ReserveCreditResponseType(str, Enum):
    credit_reserved = "credit_reserved"
    credit_exceeded = "credit_exceeded"


class ReserveCreditResponseData(BaseModel):
    order_id: int
    customer_id: int


class ReserveCreditResponse(BaseModel):
    type: ReserveCreditResponseType
    data: ReserveCreditResponseData


###


class OrderAggregateStore:
    def __init__(self):
        self.id_counter = 0
        self.orders = []

    def create_order(self, customer_id: int) -> Order:
        self.id_counter += 1

        new_order = Order(
            id=self.id_counter,
            status=OrderStatus.pending,
            customer_id=customer_id
        )
        self.orders.append(new_order.copy())
        return new_order

    def get_by_id(self, id: int) -> Optional[Order]:
        for order in self.orders:
            if order.id == id:
                return order
        return None

    def update_status_by_id(self, id: int, status: OrderStatus) -> Optional[Order]:
        for order in self.orders:
            if order.id == id:
                order.status = status
                return order
        return None


class OrderSagaOrchestrator:
    def __init__(self, order_store: OrderAggregateStore):
        self.order_store = order_store
        self._pending_credit_reserve_futures = {}

    async def create_order(self, customer_id: int) -> Order:
        pending_order = self.order_store.create_order(customer_id)
        await self._reserve_credit_for_order(pending_order)
        return self.order_store.get_by_id(pending_order.id)

    async def _reserve_credit_for_order(self, order: Order):
        result = asyncio.get_event_loop().create_future()
        self._pending_credit_reserve_futures[order.id] = result

        # Send reserve credit command
        command = CustomerCommand(
            type=CustomerCommandType.reserve_credit,
            data=CustomerCommandData(
                order_id=order.id,
                customer_id=order.customer_id
            )
        )

        async with httpx.AsyncClient() as client:
            await client.post(
                CUSTOMER_COMMAND_CHANNEL,
                json=command.dict()
            )
        
        # Wait for response to be recieved on the saga reply channel
        reserve_credit_response = await result
        print("reserve response", reserve_credit_response)
        if reserve_credit_response.type == ReserveCreditResponseType.credit_reserved:
            self.order_store.update_status_by_id(reserve_credit_response.data.order_id, OrderStatus.approved)
        elif reserve_credit_response.type == ReserveCreditResponseType.credit_exceeded:
            self.order_store.update_status_by_id(reserve_credit_response.data.order_id, OrderStatus.rejected)

    def recieved_reserve_credit_response(self, response: ReserveCreditResponse):
        future = self._pending_credit_reserve_futures.get(response.data.order_id)
        if future:
            future.set_result(response)


class OrderService:
    def __init__(self):
        self.order_store = OrderAggregateStore()
        self.create_order_orchestrator = OrderSagaOrchestrator(order_store=self.order_store)

    async def create_order(self, customer_id: int) -> Order:
        # todo: (summary of below) return pending response and handle saga in background
        # todo: make sure implementation is correct (currently it is awaiting for all response instead of -> pending order (return) -> saga handling - is it meant to be like this?)
        # todo: make seperate orchestrator for each request
        # orchestrator = OrderSagaOrchestrator(order_store=self.order_store)
        return await self.create_order_orchestrator.create_order(customer_id)


###

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    app.state.order_service = OrderService()


@app.get("/order/{order_id}")
async def fetch_order(order_id: int) -> Order:
    order = app.state.order_service.order_store.get_by_id(order_id)
    if order is not None:
        return order
    raise HTTPException(status_code=404, detail="Order Not Found")


@app.post("/order")
async def place_order(data: PlaceOrderRequest) -> Order:
    # currently returned with instant state
    order = await app.state.order_service.create_order(data.customer_id)
    return order


@app.post("/reserve_credit_response")
async def reserve_credit_response(response: ReserveCreditResponse):
    app.state.order_service.create_order_orchestrator.recieved_reserve_credit_response(response)


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5001, log_level="debug")

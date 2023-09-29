from enum import Enum
from typing import Tuple, Optional

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel


class OrderStatus(str, Enum):
    pending = "pending"
    rejected = "rejected"
    approved = "approved"


class Order(BaseModel):
    id: int
    status: OrderStatus
    item_id: int
    customer_id: int


class OrderEvent(BaseModel):
    type: str
    data: Order


class CustomerEventType(str, Enum):
    credit_reserved = "credit_reserved"
    credit_exceeded = "credit_exceeded"


class CustomerEventData(BaseModel):
    id: int
    order_id: int


class CustomerEvent(BaseModel):
    type: CustomerEventType
    data: CustomerEventData


class PlaceOrderRequest(BaseModel):
    item_id: int
    customer_id: int


class EventDispatcher:
    def __init__(self):
        self.EVENT_CHANNEL = "http://localhost:5000/order"

    async def order_created(self, order: Order) -> None:
        event = OrderEvent(
            type="created",
            data=order
        )

        async with httpx.AsyncClient() as client:
            response = await client.post(self.EVENT_CHANNEL, data=event.json())
            print(response)


app = FastAPI()
orders = []
order_counter = 0
event_dispatcher = EventDispatcher()


def get_order_by_id(id: int) -> Tuple[Optional[int], Optional[Order]]:
    for i, orders in enumerate(orders):
        if orders.id == id:
            return i, orders
    return None, None


@app.get("/order/{order_id}")
async def fetch_order(order_id: int):
    _, order = get_order_by_id(order_id)
    if order:
        return order
    raise HTTPException(status_code=404, detail="Order Not Found")


@app.post("/order")
async def place_order(data: PlaceOrderRequest):
    order_counter += 1

    created_order = Order(
        id=order_counter,
        status=OrderStatus.pending,
        item_id=data.item_id,
        customer_id=data.customer_id
    )
    orders.append(created_order)
    await event_dispatcher.order_created(created_order)

    return created_order


@app.post("/events")
async def events_handler(event: CustomerEvent):
    index, order = get_order_by_id(event.data.order_id)
    if order is None:
        raise HTTPException(status_code=404, detail="Order Not Found")

    if event.type == CustomerEventType.credit_reserved:
        # credit was succesfully reserved
        order.status = OrderStatus.approved
        orders[index] = order

        return {"status": "event processed"}
    elif event.type == CustomerEventType.credit_exceeded:
        # credit was not reserved
        order.status = OrderStatus.rejected
        orders[index] = order

        return {"status": "event processed"}
    
    raise HTTPException(status_code=422, detail="Unprocessable Event")


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5001, log_level="info")

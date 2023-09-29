from typing import Optional, Tuple
from enum import Enum

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel


class Customer(BaseModel):
    id: int
    balance: int


class OrderEventType(str, Enum):
    created = "created"
    updated = "updated"
    deleted = "deleted"


class OrderEventData(BaseModel):
    id: int
    status: str
    item_id: int
    customer_id: int


class OrderEvent(BaseModel):
    type: OrderEventType
    data: OrderEventData


class CustomerEventType(str, Enum):
    credit_reserved = "credit_reserved"
    credit_exceeded = "credit_exceeded"


class CustomerEventData(BaseModel):
    id: int
    order_id: int


class CustomerEvent(BaseModel):
    type: CustomerEventType
    data: CustomerEventData


class EventDispatcher:
    def __init__(self):
        self.EVENT_CHANNEL = "http://localhost:5000/customer"

    async def dispatch(self, event: CustomerEvent) -> None:
        async with httpx.AsyncClient() as client:
            await client.post(self.EVENT_CHANNEL, data=event.json())


app = FastAPI()
event_dispatcher = EventDispatcher()
customers = [
    Customer(id=1, balance=100),
    Customer(id=2, balance=50),
]


def get_customer_by_id(id: int) -> Tuple[Optional[int], Optional[Customer]]:
    for i, customer in enumerate(customers):
        if customer.id == id:
            return i, customer
    return None, None


@app.post("/events")
async def events_handler(event: OrderEvent) -> CustomerEvent:
    if event.type == OrderEventType.created:
        # todo: exper with item_id -> item price
        customer_index, customer = get_customer_by_id(event.data.customer_id)
        if customer is None:
            raise HTTPException(status_code=404, detail="Customer Not Found")
        
        # Attempt to reserve credit
        ITEM_PRICE = 50
        if (customer.balance - ITEM_PRICE) >= 0:
            # Credit can be reserved
            customers[customer_index] = Customer(
                id=customer.id,
                balance=(customer.balance - ITEM_PRICE)
            )

            # Credit succesfully reserved
            await event_dispatcher.dispatch(CustomerEvent(
                type=CustomerEventType.credit_reserved,
                data=CustomerEventData(
                    id=customer.id,
                    order_id=event.data.id
                )
            ))

            return {"status": "event processed"}
        else:
            # Credit cannot be reserved
            await event_dispatcher.dispatch(CustomerEvent(
                type=CustomerEventType.credit_exceeded,
                data=CustomerEventData(
                    id=customer.id,
                    order_id=event.data.id
                )
            ))

            return {"status": "event processed"}

    raise HTTPException(status_code=422, detail="Unprocessable Event")


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5002, log_level="info")

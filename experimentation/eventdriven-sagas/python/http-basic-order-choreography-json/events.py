from enum import Enum

import httpx
import uvicorn
from fastapi import FastAPI
from pydantic import BaseModel

ORDER_SERVICE_EVENT_RECIEVER = "http://localhost:5001/events"
CUSTOMER_SERVICE_EVENT_RECIEVER = "http://localhost:5002/events"

app = FastAPI()


class OrderEventType(str, Enum):
    created = "created"


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


@app.post("/order")
async def order_events_channel(event: OrderEvent) -> dict:
    async with httpx.AsyncClient() as client:
        response = await client.post(CUSTOMER_SERVICE_EVENT_RECIEVER, json=event.dict())
        print("order event routed", response.json())

    return {"status": "event routed"}


@app.post("/customer")
async def customer_events_channel(event: CustomerEvent) -> dict:
    async with httpx.AsyncClient() as client:
        response = await client.post(ORDER_SERVICE_EVENT_RECIEVER, json=event.dict())
        print("customer event routed", response.json())
    
    return {"status": "event routed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000, log_level="debug")

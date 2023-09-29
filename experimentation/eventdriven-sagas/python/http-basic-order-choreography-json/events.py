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


class CustomerEvent(BaseModel):
    customer_id: str


@app.post("/order")
async def order_events_channel(event: OrderEvent) -> dict:
    async with httpx.AsyncClient() as client:
        response = await client.post(CUSTOMER_SERVICE_EVENT_RECIEVER, data=event.json())
        print("order event routed", response)

    return {"status": "event routed"}


@app.post("/customer")
async def customer_events_channel(event: CustomerEvent) -> dict:
    async with httpx.AsyncClient() as client:
        response = await client.post(ORDER_SERVICE_EVENT_RECIEVER, data=event.json())
        print("customer event routed", response)
    
    return {"status": "event routed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000, log_level="info")

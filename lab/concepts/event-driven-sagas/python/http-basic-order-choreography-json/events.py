from enum import Enum

import httpx
import uvicorn
from fastapi import FastAPI
from pydantic import BaseModel

ORDER_SERVICE_EVENT_RECIEVER = "http://localhost:5001/events"
CUSTOMER_SERVICE_EVENT_RECIEVER = "http://localhost:5002/events"

app = FastAPI()


@app.post("/order")
async def order_events_channel(event: dict) -> dict:
    print("order event", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(CUSTOMER_SERVICE_EVENT_RECIEVER, json=event)
        print("order event - routed", response.json())

    return {"status": "event routed"}


@app.post("/customer")
async def customer_events_channel(event: dict) -> dict:
    print("customer event", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(ORDER_SERVICE_EVENT_RECIEVER, json=event)
        print("customer event routed", response.json())
    
    return {"status": "event routed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000, log_level="debug")

import httpx
import uvicorn
from fastapi import FastAPI

ORDER_SERVICE_EVENT_RECIEVER_1 = "http://localhost:5001/customer_event_consumer"
ORDER_SERVICE_EVENT_RECIEVER_2 = "http://localhost:5001/inventory_event_consumer"
CUSTOMER_SERVICE_EVENT_RECIEVER_1 = "http://localhost:5002/order_event_consumer"
INVENTORY_SERVICE_EVENT_RECIEVER_1 = "http://localhost:5003/order_event_consumer"

app = FastAPI()


@app.post("/order_event")
async def order_event_channel(event: dict) -> dict:
    print("order event", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(CUSTOMER_SERVICE_EVENT_RECIEVER_1, json=event)
        print("order event - routed to customer", response.json())

    async with httpx.AsyncClient() as client:
        response = await client.post(INVENTORY_SERVICE_EVENT_RECIEVER_1, json=event)
        print("order event - routed to inventory", response.json())

    return {"status": "event routed"}


@app.post("/customer_event")
async def customer_event_channel(event: dict) -> dict:
    print("customer event", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(ORDER_SERVICE_EVENT_RECIEVER_1, json=event)
        print("customer event - routed to order", response.json())
    
    return {"status": "event routed"}


@app.post("/inventory_event")
async def inventory_event_channel(event: dict) -> dict:
    print("inventory event", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(ORDER_SERVICE_EVENT_RECIEVER_2, json=event)
        print("inventory event - routed to order", response.json())
    
    return {"status": "event routed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000, log_level="debug")

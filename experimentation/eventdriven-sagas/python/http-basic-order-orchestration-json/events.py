import httpx
import uvicorn
from fastapi import FastAPI

CREATE_ORDER_SAGA_REPLY_RECIEVER = "http://localhost:5001/reserve_credit_response"
CUSTOMER_COMMAND_CHANNEL_RECIEVER = "http://localhost:5002/customer_commands"

app = FastAPI()


@app.post("/customer_commands")
async def customer_command_channel(event: dict) -> dict:
    print("customer command", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(CUSTOMER_COMMAND_CHANNEL_RECIEVER, json=event)

    return {"status": "event routed"}


@app.post("/create_order_reply")
async def create_order_reply_channel(event: dict) -> dict:
    print("create order reply", event)
    async with httpx.AsyncClient() as client:
        response = await client.post(CREATE_ORDER_SAGA_REPLY_RECIEVER, json=event)
    
    return {"status": "event routed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000, log_level="debug")

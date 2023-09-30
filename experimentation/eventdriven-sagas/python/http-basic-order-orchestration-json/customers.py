from enum import Enum
from typing import Tuple, Optional

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel


CREATE_ORDER_REPLY_CHANNEL = "http://localhost:5000/create_order_reply"


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


class Customer(BaseModel):
    id: int
    balance: int


###


class CustomerService:
    def __init__(self):
        self.customers = [
            Customer(id=1, balance=100),
            Customer(id=2, balance=50),
        ]

    def _get_customer_by_id(self, id: int) -> Tuple[Optional[int], Optional[Customer]]:
        for i, customer in enumerate(self.customers):
            if customer.id == id:
                return i, customer
            
        return None, None
    
    async def handle_command(self, command: CustomerCommand):
        customer_index, customer = self._get_customer_by_id(command.data.customer_id)
        if customer is None:
            raise HTTPException(status_code=404, detail="Customer Not Found")

        # Attempt to reserve credit
        ITEM_PRICE = 50
        if (customer.balance - ITEM_PRICE) >= 0:
            # Credit can be reserved
            self.customers[customer_index] = Customer(
                id=customer.id,
                balance=(customer.balance - ITEM_PRICE)
            )

            # Credit succesfully reserved
            response = ReserveCreditResponse(
                type=ReserveCreditResponseType.credit_reserved,
                data=ReserveCreditResponseData(
                    order_id=command.data.order_id,
                    customer_id=customer.id
                )
            )

            async with httpx.AsyncClient() as client:
                await client.post(
                    CREATE_ORDER_REPLY_CHANNEL,
                    json=response.dict()
                )
        else:
            # Credit cannot be reserved
            response = ReserveCreditResponse(
                type=ReserveCreditResponseType.credit_exceeded,
                data=ReserveCreditResponseData(
                    order_id=command.data.order_id,
                    customer_id=customer.id
                )
            )
            async with httpx.AsyncClient() as client:
                await client.post(
                    CREATE_ORDER_REPLY_CHANNEL,
                    json=response.dict()
                )


###

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    app.state.customer_service = CustomerService()


@app.post("/customer_commands")
async def commands_channel(command: CustomerCommand):
    await app.state.customer_service.handle_command(command)
    return {"detail": "command processed"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5002, log_level="debug")

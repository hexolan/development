from typing import List

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException

from test_schema.schema_models import Customer
from test_schema.schema_events import OrderEvent, OrderEventType, OrderRejectionCause, CustomerCreditEvent, CustomerCreditEventType, CustomerCreditEventData

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    app.state.customers: List[Customer] = [
        Customer(id=1, balance=100),
        Customer(id=2, balance=50),
    ]


@app.get("/customers_debug")
def view_customers():
    return app.state.customers


@app.post("/order_event_consumer")
async def order_event_consumer(event: OrderEvent):
    # constants
    ITEM_PRICE = 50

    # find customer
    index, customer = None, None
    for i, cust in enumerate(app.state.customers):
        if cust.id == event.data.customer_id:
            index, customer = i, cust

    if not customer:
        raise HTTPException(status_code=404, detail="Customer Not Found")

    # if a pending order has been created
    if event.type == OrderEventType.created:
        # check balance after cost of purchase
        if (customer.balance - ITEM_PRICE) >= 0:
            # Balance can be reserved
            customer.balance -= ITEM_PRICE
            app.state.customers[index] = customer

            # Dispatch event
            async with httpx.AsyncClient() as client:
                await client.post(
                    "http://localhost:5000/customer_event",
                    json=CustomerCreditEvent(
                        type=CustomerCreditEventType.credit_reserved,
                        data=CustomerCreditEventData(
                            id=customer.id,
                            order_id=event.data.id
                        )
                    ).dict()
                )
        else:
            # Credit exceeded
            # Dispatch event
            async with httpx.AsyncClient() as client:
                await client.post(
                    "http://localhost:5000/customer_event",
                    json=CustomerCreditEvent(
                        type=CustomerCreditEventType.credit_exceeded,
                        data=CustomerCreditEventData(
                            id=customer.id,
                            order_id=event.data.id
                        )
                    ).dict()
                )
    elif event.type == OrderEventType.rejected:
        # make sure rejection wasn't as a result of insufficient balance (otherwise i'm creating customer balance)
        if event.detail != OrderRejectionCause.customer_balance:
            # rollback reserved balance as order was rejected
            print("credit reservation rollback")
            customer.balance += ITEM_PRICE
            app.state.customers[index] = customer


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5002, log_level="debug")

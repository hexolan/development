from typing import List

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException

from test_schema.schema_models import Order, OrderStatus, PlaceOrderRequest
from test_schema.schema_events import (
    OrderEvent, OrderEventType, OrderRejectionCause,
    CustomerCreditEvent, CustomerCreditEventType,
    InventoryReservationEvent, InventoryReservationEventType
)

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    app.state.order_count = int(0)
    app.state.orders: List[Order] = []
    app.state.pending_order_state = {}


@app.get("/order/{order_id}")
async def view_order(order_id: int) -> Order:
    # Get order
    for order in app.state.orders:
        if order.id == order_id:
            return order
    
    raise HTTPException(status_code=404, detail="Order Not Found")


@app.post("/order")
async def place_order(data: PlaceOrderRequest) -> Order:
    # create pending order
    app.state.order_count += 1
    pending_order = Order(
        id=app.state.order_count,
        status=OrderStatus.pending,
        customer_id=data.customer_id,
        product_id=data.product_id
    )

    # state of order (e.g. inventory response recieved - indicative of sufficient stock)
    app.state.pending_order_state[pending_order.id] = {
        "inventory_approved": None,
        "customer_approved": None
    }

    app.state.orders.append(pending_order.copy())

    # Dispatch event
    async with httpx.AsyncClient() as client:
        await client.post(
            "http://localhost:5000/order_event",
            json=OrderEvent(
                type=OrderEventType.created,
                data=pending_order
            ).dict()
        )

    return pending_order


@app.post("/customer_event_consumer")
async def customer_event_consumer(event: CustomerCreditEvent):
    # todo: logic wise
    
    # idea:
    # rejection response => instant order rejection
    # but then aggregating all the acceptances? some state

    # payment approved && item in stock booleans on order?
    # or pending order table (and upon reciept of updates check against table)
    # then when all is true order can be approved on reciept of final required event?

    # current approach: state records for pending orders
    pending_state = app.state.pending_order_state.get(event.data.order_id)
    if pending_state:
        if event.type == CustomerCreditEventType.credit_reserved:
            # approved by customer service
            pending_state["customer_approved"] = True
        elif event.type == CustomerCreditEventType.credit_exceeded:
            # customer does not have sufficient balance
            pending_state["customer_approved"] = False
        
        app.state.pending_order_state[event.data.order_id] = pending_state

        # check pending order state (to see if order is ready to be approved / should be rejected)
        await review_order(event.data.order_id)


@app.post("/inventory_event_consumer")
async def inventory_event_consumer(event: InventoryReservationEvent):
    pending_state = app.state.pending_order_state.get(event.data.order_id)
    if pending_state:
        if event.type == InventoryReservationEventType.reserved_stock:
            # approved by inventory service
            pending_state["inventory_approved"] = True
        elif event.type == InventoryReservationEventType.out_of_stock:
            # not enough stock of item
            pending_state["inventory_approved"] = False

        app.state.pending_order_state[event.data.order_id] = pending_state

        # check pending order state (to see if order is ready to be approved / should be rejected)
        await review_order(event.data.order_id)


async def review_order(id: int):
    # get order
    index, order = None, None
    for i, orderSearch in enumerate(app.state.orders):
        if orderSearch.id == id:
            index, order = i, orderSearch

    # check if order should be approved/rejected
    approve_order = True
    reject_order = False

    pending_state = app.state.pending_order_state.get(id)
    for _, value in pending_state.items():
        if value == False:
            reject_order = True
            approve_order = False
            break
        elif value == None:
            approve_order = False

    # approve/rejection logic
    if approve_order:
        # Mark order as approved
        order.status = OrderStatus.approved
        app.state.orders[index] = order
        del app.state.pending_order_state[order.id]

        # Dispatch approved event
        async with httpx.AsyncClient() as client:
            await client.post(
                "http://localhost:5000/order_event",
                json=OrderEvent(
                    type=OrderEventType.approved,
                    data=order
                ).dict()
            )
    elif reject_order:
        # Mark order as rejected
        order.status = OrderStatus.rejected
        app.state.orders[index] = order
        del app.state.pending_order_state[order.id]

        # Dispatch rejected event
        # todo: add rejection detail (caused by stock? caused by customer?)
        #   > also think about cases where both stock and credit resulted in failure (but only one processed before rejection sent out)
        #       > this is currently the case with inventory and customer (when both fail)
        #       * rollback is performed before inventory service has even processed evnet
        #       * then when processed it uses the 'created' inventory to create appearence of having reserved inventory
        #       * also making it look like no issue occured as the resultant inventory is the same as before order rejection
        #   (in short - this current version is not safe for real world use)
        #   > as part of rollback inventory should not be created

        rejection_cause = None
        if pending_state["inventory_approved"] == False:
            rejection_cause = OrderRejectionCause.inventory_stock
        elif pending_state["customer_approved"] == False:
            rejection_cause = OrderRejectionCause.customer_balance

        async with httpx.AsyncClient() as client:
            await client.post(
                "http://localhost:5000/order_event",
                json=OrderEvent(
                    type=OrderEventType.rejected,
                    data=order,
                    detail=rejection_cause
                ).dict()
            )
    else:
        # still waiting on more results
        print("order still pending")


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5001, log_level="debug")

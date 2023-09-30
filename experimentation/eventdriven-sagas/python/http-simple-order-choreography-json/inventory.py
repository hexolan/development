from typing import List

import httpx
import uvicorn
from fastapi import FastAPI, HTTPException

from test_schema.schema_models import Inventory
from test_schema.schema_events import OrderEvent, OrderEventType, OrderRejectionCause, InventoryReservationEvent, InventoryReservationEventType, InventoryReservationEventData

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    app.state.inventory: List[Inventory] = [
        Inventory(product_id=1, quantity=2),
        Inventory(product_id=2, quantity=0),
    ]


@app.post("/order_event_consumer")
async def order_event_consumer(event: OrderEvent):
    # find product inventory
    index, product_inv = None, None
    for i, inventory in enumerate(app.state.inventory):
        if inventory.product_id == event.data.product_id:
            index, product_inv = i, inventory
    
    if not product_inv:
        raise HTTPException(status_code=404, detail="Inventory Not Found")


    # if a pending order has been created
    if event.type == OrderEventType.created:
        # check quantity after reservation (out of stock or stock reserved)
        if (product_inv.quantity - 1) >= 0:
            # Item can be reserved
            product_inv.quantity -= 1
            app.state.inventory[index] = product_inv

            # Dispatch event
            async with httpx.AsyncClient() as client:
                await client.post(
                    "http://localhost:5000/inventory_event",
                    json=InventoryReservationEvent(
                        type=InventoryReservationEventType.reserved_stock,
                        data=InventoryReservationEventData(
                            product_id=product_inv.product_id,
                            order_id=event.data.id
                        ),
                    ).dict()
                )
        else:
            # Item out of stock
            # Dispatch event
            async with httpx.AsyncClient() as client:
                await client.post(
                    "http://localhost:5000/inventory_event",
                    json=InventoryReservationEvent(
                        type=InventoryReservationEventType.out_of_stock,
                        data=InventoryReservationEventData(
                            product_id=product_inv.product_id,
                            order_id=event.data.id
                        ),
                    ).dict()
                )
    elif event.type == OrderEventType.rejected:
        # make sure rejection wasn't as a result of insufficient stock (otherwise i'm creating stock)
        if event.detail != OrderRejectionCause.inventory_stock:
            # rollback stock reservation as order was rejected
            product_inv.quantity += 1
            app.state.inventory[index] = product_inv


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5003, log_level="debug")

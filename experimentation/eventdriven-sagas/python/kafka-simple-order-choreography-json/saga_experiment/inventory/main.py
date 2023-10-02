import asyncio
from json import loads
from typing import List

import uvicorn
from fastapi import FastAPI
from aiokafka import AIOKafkaConsumer, AIOKafkaProducer

from saga_experiment.utils.schema_models import Inventory
from saga_experiment.utils.schema_events import (
    InventoryReservationEvent, InventoryReservationEventData, InventoryReservationEventType,
    CustomerCreditEvent, CustomerCreditEventType
)


class InventoryConsumer:
    def __init__(self, app: FastAPI) -> None:
        self.app = app

    async def start(self) -> None:
        # create consumer and producer
        KAFKA_BOOTSTRAP_SERVERS = "localhost:9092"
        self.consumer = AIOKafkaConsumer("customer", group_id="inventory-service", bootstrap_servers=KAFKA_BOOTSTRAP_SERVERS)
        self.producer = AIOKafkaProducer(bootstrap_servers=KAFKA_BOOTSTRAP_SERVERS)

        # start connections
        await self.producer.start()
        await self.consumer.start()
  
        # begin consuming
        try:
            await self.begin_consuming()
        finally:
            await self.producer.stop()
            await self.consumer.stop()
    
    async def begin_consuming(self) -> None:
        async for msg in self.consumer:
            await self.process_msg(msg)

    async def process_msg(self, msg) -> None:
        # consuming customer events - upon reciept of credit reserved, perform inv reservation
        # print("consumed: ", msg.topic, msg.partition, msg.offset, msg.key, msg.value, msg.timestamp)
        event = CustomerCreditEvent(**loads(msg.value.decode("utf-8")))
        await self.process_customer_event(event)

    async def process_customer_event(self, event: CustomerCreditEvent):
        # find product inventory
        index, product_inv = None, None
        for i, inventory in enumerate(self.app.state.inventory):
            if inventory.product_id == event.data.product_id:
                index, product_inv = i, inventory

        if event.type == CustomerCreditEventType.credit_reserved:
            # check quantity after reservation (out of stock or stock reserved)
            if (product_inv.quantity - 1) >= 0:
                # Item can be reserved
                product_inv.quantity -= 1
                self.app.state.inventory[index] = product_inv

                # Dispatch event
                await self.dispatch_reserved(
                    product_id=event.data.product_id,
                    order_id=event.data.order_id,
                    customer_id=event.data.customer_id
                )
            else:
                # not enough stock of item
                # Dispatch event
                await self.dispatch_out_of_stock(
                    product_id=event.data.product_id,
                    order_id=event.data.order_id,
                    customer_id=event.data.customer_id
                )

    async def dispatch_reserved(self, product_id: int, order_id: int, customer_id: int) -> None:
        event = InventoryReservationEvent(
            type=InventoryReservationEventType.reserved_stock,
            data=InventoryReservationEventData(
                product_id=product_id,
                order_id=order_id,
                customer_id=customer_id
            )
        )
        event_as_bytes = event.model_dump_json().encode("utf-8")
        await self.producer.send_and_wait("inventory", event_as_bytes)

    async def dispatch_out_of_stock(self, product_id: int, order_id: int, customer_id: int) -> None:
        event = InventoryReservationEvent(
            type=InventoryReservationEventType.out_of_stock,
            data=InventoryReservationEventData(
                product_id=product_id,
                order_id=order_id,
                customer_id=customer_id
            )
        )
        event_as_bytes = event.model_dump_json().encode("utf-8")
        await self.producer.send_and_wait("inventory", event_as_bytes)



app = FastAPI()
consumer_app = InventoryConsumer(app)


@app.on_event("startup")
async def startup_event():
    app.state.inventory: List[Inventory] = [
        Inventory(product_id=1, quantity=2),
        Inventory(product_id=2, quantity=0),
    ]
    asyncio.create_task(consumer_app.start())


def start() -> None:
    uvicorn.run(app, host="0.0.0.0", port=5003, log_level="debug")


if __name__ == "__main__":
    start()
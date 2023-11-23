import asyncio
from json import loads
from typing import List

import uvicorn
from fastapi import FastAPI
from aiokafka import AIOKafkaConsumer, AIOKafkaProducer

from saga_experiment.utils.schema_models import Customer
from saga_experiment.utils.schema_events import (
    CustomerCreditEvent, CustomerCreditEventData, CustomerCreditEventType,
    OrderEvent, OrderEventType,
    InventoryReservationEvent, InventoryReservationEventType
)


class CustomerConsumer:
    def __init__(self, app: FastAPI) -> None:
        self.app = app

    async def start(self) -> None:
        # create consumer and producer
        KAFKA_BOOTSTRAP_SERVERS = "localhost:9092"
        self.consumer = AIOKafkaConsumer("order", "inventory", group_id="customer-service", bootstrap_servers=KAFKA_BOOTSTRAP_SERVERS)
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
        # consuming both order and inventory events
        # on reciept of order event
        #   - attempt to reserve customer credit for that order
        #   - dispatch own (customer service event) of result
        # on reciept of inventory event
        #   - if the inventory event is a inventory reservation event
        #   - .. and the event indicates failure
        #   - ... then perform a rollback of credit reserved for the user
        async for msg in self.consumer:
            await self.process_msg(msg)

    async def process_msg(self, msg) -> None:
        # print("consumed: ", msg.topic, msg.partition, msg.offset, msg.key, msg.value, msg.timestamp)
        value_as_str = msg.value.decode("utf-8")
        if msg.topic == "order":
            event = OrderEvent(**loads(value_as_str))
            await self.process_order_event(event)
        elif msg.topic == "inventory":
            event = InventoryReservationEvent(**loads(value_as_str))
            await self.process_inventory_event(event)

    async def process_order_event(self, event: OrderEvent):
        # constants
        ITEM_PRICE = 50

        # find customer
        index, customer = None, None
        for i, cust in enumerate(self.app.state.customers):
            if cust.id == event.data.customer_id:
                index, customer = i, cust

        if event.type == OrderEventType.created:
            # attempt to reserve credit for customer
            # check balance after cost of purchase
            if (customer.balance - ITEM_PRICE) >= 0:
                # Balance can be reserved
                customer.balance -= ITEM_PRICE
                self.app.state.customers[index] = customer

                # Dispatch event (succesful)
                await self.dispatch_reserved(customer_id=customer.id, order_id=event.data.id, product_id=event.data.product_id)
            else:
                # Customer does not have enough balance
                await self.dispatch_exceeded(customer_id=customer.id, order_id=event.data.id, product_id=event.data.product_id)

    async def process_inventory_event(self, event: InventoryReservationEvent):
        # constants
        ITEM_PRICE = 50

        # find customer
        index, customer = None, None
        for i, cust in enumerate(self.app.state.customers):
            if cust.id == event.data.customer_id:
                index, customer = i, cust

        if event.type == InventoryReservationEventType.out_of_stock:
            # perform compensation transaction (give user their credit back)
            customer.balance += ITEM_PRICE
            self.app.state.customers[index] = customer
            # todo: potential customer balance update event?

    async def dispatch_reserved(self, customer_id: int, order_id: int, product_id: int) -> None:
        event = CustomerCreditEvent(
            type=CustomerCreditEventType.credit_reserved,
            data=CustomerCreditEventData(
                customer_id=customer_id,
                order_id=order_id,
                product_id=product_id
            )
        )
        event_as_bytes = event.model_dump_json().encode("utf-8")
        await self.producer.send_and_wait("customer", event_as_bytes)

    async def dispatch_exceeded(self, customer_id: int, order_id: int, product_id: int) -> None:
        event = CustomerCreditEvent(
            type=CustomerCreditEventType.credit_exceeded,
            data=CustomerCreditEventData(
                customer_id=customer_id,
                order_id=order_id,
                product_id=product_id
            )
        )
        event_as_bytes = event.model_dump_json().encode("utf-8")
        await self.producer.send_and_wait("customer", event_as_bytes)


app = FastAPI()
consumer_app = CustomerConsumer(app)


@app.on_event("startup")
async def startup_event():
    app.state.customers: List[Customer] = [
        Customer(id=1, balance=100),
        Customer(id=2, balance=50),
    ]
    asyncio.create_task(consumer_app.start())


def start() -> None:
    uvicorn.run(app, host="0.0.0.0", port=5002, log_level="debug")


if __name__ == "__main__":
    start()
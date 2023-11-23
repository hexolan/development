import asyncio
from json import loads
from typing import List

import uvicorn
from fastapi import FastAPI, HTTPException
from aiokafka import AIOKafkaConsumer, AIOKafkaProducer

from saga_experiment.utils.schema_models import Order, OrderStatus, PlaceOrderRequest
from saga_experiment.utils.schema_events import (
    OrderEvent, OrderEventType,
    InventoryReservationEvent, InventoryReservationEventType,
    CustomerCreditEvent, CustomerCreditEventType
)


class OrderConsumer:
    def __init__(self, app: FastAPI) -> None:
        self.app = app

    async def start(self) -> None:
        # create consumer and producer
        KAFKA_BOOTSTRAP_SERVERS = "localhost:9092"
        self.consumer = AIOKafkaConsumer("inventory", "customer", group_id="order-service", bootstrap_servers=KAFKA_BOOTSTRAP_SERVERS)
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
        # print("consumed: ", msg.topic, msg.partition, msg.offset, msg.key, msg.value, msg.timestamp)
        value_as_str = msg.value.decode("utf-8")
        if msg.topic == "inventory":
            event = InventoryReservationEvent(**loads(value_as_str))
            await self.process_inventory_event(event)
        elif msg.topic == "customer":
            event = CustomerCreditEvent(**loads(value_as_str))
            await self.process_customer_event(event)

    async def process_inventory_event(self, event: InventoryReservationEvent) -> None:
        # fetch order
        index, order = None, None
        for i, orderSearch in enumerate(self.app.state.orders):
            if orderSearch.id == event.data.order_id:
                index, order = i, orderSearch

        if event.type == InventoryReservationEventType.reserved_stock:
            print("good - stock reserved")
            # as inventory will be checked after customer credit - if this is succesful
            # then the order can be approved.

            # mark order as approved
            order.status = OrderStatus.approved
            self.app.state.orders[index] = order

            # todo: dispatch approved (/ status update) event
        elif event.type == InventoryReservationEventType.out_of_stock:
            print("bad - item out of stock")

            # mark order as rejected
            order.status = OrderStatus.rejected
            self.app.state.orders[index] = order

            # todo: dispatch rejected (/ status update) event

    async def process_customer_event(self, event: CustomerCreditEvent) -> None:
        # fetch order
        index, order = None, None
        for i, orderSearch in enumerate(self.app.state.orders):
            if orderSearch.id == event.data.order_id:
                index, order = i, orderSearch
    
        if event.type == CustomerCreditEventType.credit_reserved:
            print("good - credit succesfully reserved")
            # now wait until a out of stock event is recieved back
        elif event.type == CustomerCreditEventType.credit_exceeded:
            print("bad - credit exceeded")

            # mark order as rejected
            order.status = OrderStatus.rejected
            self.app.state.orders[index] = order

            # todo: dispatch rejected (/ status update) event

    async def dispatch_order_created(self, order: Order) -> None:
        event = OrderEvent(type=OrderEventType.created, data=order)
        event_as_bytes = event.model_dump_json().encode("utf-8")

        """
        Notes:
        planned order
        Order Created Event
        -> Customer Service (attempts to reserve credit)
            -> Credit Reserved Event

        -> Credit Reserved Event consumed by inv service (only if succesful)
            - Inventory service attempts to reserve stock
            -> Inventory Reservation Event

        -> Customer Service listens to inv reservation event
            -> if failure then customer service will rollback reservation of credit

        -> Order Services listens to
            -> customer service events (to mark order rejected)
            -> inventory service events (to mark order rejected - or approved if stock succesful)
        """

        await self.producer.send_and_wait("order", event_as_bytes)



app = FastAPI()
consumer_app = OrderConsumer(app)


@app.on_event("startup")
async def startup_event():
    app.state.id_count = 0
    app.state.orders: List[Order] = []
    asyncio.create_task(consumer_app.start())


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
    app.state.id_count += 1
    pending_order = Order(
        id=app.state.id_count,
        status=OrderStatus.pending,
        customer_id=data.customer_id,
        product_id=data.product_id
    )
    app.state.orders.append(pending_order.model_copy())

    # Dispatch created event (important that this actually gets dispatched)
    await consumer_app.dispatch_order_created(pending_order)

    # give pending order to user
    return pending_order


def start() -> None:
    uvicorn.run(app, host="0.0.0.0", port=5001, log_level="debug")


if __name__ == "__main__":
    start()
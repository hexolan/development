"""Events"""
from typing import Optional
from enum import Enum

from pydantic import BaseModel

from saga_experiment.utils.schema_models import Order


## Order Events ##


class OrderEventType(str, Enum):
    created = "created"
    approved = "approved"
    rejected = "rejected"


class OrderRejectionCause(str, Enum):
    inventory_stock = "inventory_stock"
    customer_balance = "customer_balance"


class OrderEvent(BaseModel):
    type: OrderEventType
    data: Order


## End Order Events ##


## Customer Events ##


class CustomerCreditEventType(str, Enum):
    credit_reserved = "credit_reserved"
    credit_exceeded = "credit_exceeded"


class CustomerCreditEventData(BaseModel):
    product_id: int
    order_id: int
    customer_id: int


class CustomerCreditEvent(BaseModel):
    type: CustomerCreditEventType
    data: CustomerCreditEventData


## End Customer Events ##


## Inventory Events ##


class InventoryReservationEventType(str, Enum):
    reserved_stock = "reserved_stock"
    out_of_stock = "out_of_stock"


class InventoryReservationEventData(BaseModel):
    product_id: int
    order_id: int
    customer_id: int


class InventoryReservationEvent(BaseModel):
    type: InventoryReservationEventType
    data: InventoryReservationEventData


## End Inventory Events ##

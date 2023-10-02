"""Models"""
from enum import Enum

from pydantic import BaseModel


## Orders ##


class OrderStatus(str, Enum):
    pending = "pending"
    rejected = "rejected"
    approved = "approved"


class Order(BaseModel):
    id: int
    status: OrderStatus
    customer_id: int
    product_id: int


class PlaceOrderRequest(BaseModel):
    customer_id: int
    product_id: int


## End Orders ##


## Customers ##


class Customer(BaseModel):
    id: int
    balance: int


## End Customers ##


## Inventory ##


class Inventory(BaseModel):
    product_id: int
    quantity: int


## End Inventory ##

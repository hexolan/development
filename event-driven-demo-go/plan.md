# E-Commerce Application

**Features**

Branches (each location)

Products (items served by the chain)

Product Stock (stock of each product at each branch)

- Item served at branch?
- Quantity remaining

Orders

Notifications (sending email on order updates, etc)

Payment

Users



**Planned Flows**

<u>Place Order</u>

* User places order for items in their cart (at their selected branch)
* Order created in pending state -> *OrderCreatedEvent*
* *OrderCreatedEvent* -> check stock of items in order 
  * -> OrderStockEvent (type = reserved / out_of_stock)
* *OrderStockEvent* -> reserve credit from user for order
  * -> OrderPaymentEvent (type = success / failure)
* *OrderPaymentEvent* (type: success) -> Order state changed to approved
* *OrderPaymentEvent* (type: failure) -> Stock returned & order rejected
* *OrderStockEvent* (type: out_of_stock) -> Order rejected

 
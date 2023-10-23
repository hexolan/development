# E-Commerce Application

**Features**

Reviews/Rating Service

Categories

Product

* Items served

Warehouse

- Stock keeping of each product

Orders

* Placed orders

Shipping

* Responsible for dispatching items (tracking ids for imaginary mail, etc)

Notifications (sending email on order updates, etc)

Payment

* Handling payments. User credit, etc

Users

* User profiles and details



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

* OrderShippingEvent (type: shipping_scheduled or no_avaliability)

topics:

order.created
> OrderCreateEvent

topics for creation sagas:
order.created.stock (OrderStockEvent)
order.created.payment (OrderPaymentEvent)
order.created.shipping (OrderShippingEvent)

stock.added
stock.removed

payment.debited (taken)
payment.credited (returned)

----

# Old

Orders:
> Order Created Event
> Order Updated Event
> Order Cancelled? Event

Customers:
> Customer Created Event
> Customer Updated Event
> Customer Deleted Event

-- Order Created Event leads to...
--- Credit Reservation Event
      > enum type on event (CREDIT_RESERVED, CREDIT_DENIED)

-- Inventory Reservation Event leads to..
--- (if inventory reservation event type == INVENTORY_DENIED) perform compensation transaction
      > dispatch Credit Reservation Event (with type CREDIT_RETURNED) or something

Inventory:
> Inventory Event
  .. reflects changes to inventory

-- Credit Reservation Event leads to..
--- Inventory Reservation Event
      > enum type on event (INVENTORY_RESERVED, INVENTORY_DENIED)

Although unused implement compensation capability (e.g. INVENTORY_RETURNED) type on reservation event
	> actually can be used when an order is cancelled
	> if order is pending... wait for order to be approved (or denied) before attempting to cancel an order (so the saga is not still in progress - since this is choreography not orchestrated where it can be easily interrupted)

# Topics

* order.created
* order.updated
* order.deleted

* order.creation.payment
* order.creation.warehouse
* order.creation.shipping

/order (POST) Request ->
  Order Created (dispatch OrderCreated event)
  Return Order ID to User

Order Service
  Primary:
  > Consume OrderShippingEvent
  > of type 'approved'
  > Set order status to approved

  Compensationary:
  > Consume OrderPaymentEvent, OrderWarehouseEvent
  > Set order status to rejected

Payment Service
  Primary:
  > Consume OrderCreatedEvent
  > Dispatch OrderPaymentEvent
    > Type: 'approved' or 'denied'

  Compensationary:
  > Consume OrderWarehouseEvent
    > of type 'out_of_stock'
  > Return Credit to User

Warehouse Service
  Primary:
  > Consume OrderPaymentEvent
  > Dispatch OrderWarehouseEvent
    > Type: 'reserved' or 'out_of_stock'
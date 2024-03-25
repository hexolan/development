# Event Docs: OrderStateEvent

## Overview

**Description**

Dispatched upon create, update and delete actions to ``order`` entities.

Used for implementing the [Event Sourcing](https://microservices.io/patterns/data/event-sourcing.html) pattern.

**Schemas** 

* [``/schema/protobufs/order/v1/events.proto``](/schema/protobufs/order/v1/events.proto)

**Topics**

* ``order.state.created``
  * event.type = TYPE_CREATED
* ``order.state.updated``
  * event.type = TYPE_UPDATED
* ``order.state.deleted``
  * event.type = TYPE_DELETED

## Producers and Consumers

### Producers

This event is only dispatched by the order service.

### Consumers

todo: complete documentation

## Additional Notes

n/a
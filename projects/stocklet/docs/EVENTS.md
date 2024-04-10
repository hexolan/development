# Stocklet: Event Docs

## Overview

The events are schemed and serialised using [protocol buffers](https://protobuf.dev/). The events schemas can be found in [``/schema/protobufs/stocklet/events/``](/schema/protobufs/stocklet/events/)

They are dispatched using the [transactional outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html). Debezium is used as a relay to publish events from database outbox tables to the message broker (Kafka). The Debezium connectors are configured by the ``service-init`` containers, which are also responsible for performing database migrations for their respective services.

## Services

### Auth Service

**Produces:**

* n/a

**Consumes:**

* UserDeletedEvent

### Order Service

**Produces:**

* todo:

**Consumes:**

* ProductPriceQuoteEvent
* StockReservationEvent
* ShipmentAllocationEvent
* PaymentProcessedEvent

### Payment Service

**Produces:**

* todo:

**Consumes:**

* UserCreatedEvent
* UserDeletedEvent
* ShipmentAllocationEvent

### Product Service

**Produces:**

* todo:

**Consumes:**

* OrderCreatedEvent

### Shipping Service

**Produces:**

* todo:

**Consumes:**

* StockReservationEvent
* PaymentProcessedEvent

### User Service

**Produces:**

* todo:

**Consumes:**

* n/a

### Warehouse Service

**Produces:**

* todo:

**Consumes:**

* OrderPendingEvent
* ShipmentAllocationEvent
* PaymentProcessedEvent

## Miscellaneous

### Place Order Saga

todo: documentation
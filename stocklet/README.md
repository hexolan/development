# Stocklet

An event-driven microservice e-commerce application.

TODO: add support for using NATS (/ NATS + Liftbridge) for messaging instead of Kafka

* <https://docs.nats.io/>
* <https://liftbridge.io/docs/quick-start.html>

todo: Apache -> AGPL

## About

### Layout

current state of README: rough outline

... project organised as a monorepo

[diagram of project layout .... inventory service / order service -> kafka as 'event bus' ... service mesh]

### Schema

(maybe add additional services such as shipping service -> showing eventual consistency)
  > in that case rename this 'project' to event-driven-go-demo (can also create an event-driven-py-demo, etc..)
  > might be a better angle (+ add more services // warehouse service // dispatch/delivery service // mark order delivered, etc)
  > database access services? (-> gRPC)
    > requesting coalescing

description of events (and invoked compensation transactions) + topics they are dispatched under
  > e.g. 'order.created', 'order.updated' topics

discussion on protobuf schema of events
  > potentially write event schema with <https://www.asyncapi.com/> (find way to integrate protobuf as it appears json-oriented)

## Features

technologies:
> kafka
> postgresql
> (TODO) open telemetry
> grpc
> grpc gateway
> protobuf
> NATS support
> "pluggable architecture"

maybe: description of patterns used

event-driven architecture

description of choreography saga
  > links to resources such as microservices.io

difference between choreography and orchestration

* [ ] ABC
* [ ] DEF

## Deployment

how to run project
  > write docker compose file (add kafka and kafka ui)

tracing method?
  > some sort of request ID/transaction ID that spans all saga related events (though not needed - tracing could be performed using just order ID)
  > <https://medium.com/dzerolabs/observability-journey-understanding-logs-events-traces-and-spans-836524d63172>

## License

This project is distributed under the [XYZ License.](/LICENSE)

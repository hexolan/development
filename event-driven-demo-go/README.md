rough outline/version

# Event Driven Demonstration (Golang)

## Layout

[diagram of project layout .... inventory service / order service -> kafka as 'event bus']

## Events

description of events (and invoked compensation transactions) + topics they are dispatched under
  > e.g. 'order.created', 'order.updated' topics

## Patterns Used

event-driven architecture

description of choreography saga
  > links to resources such as microservices.io

difference between choreography and orchestration

## Schema

(maybe add additional services such as shipping service -> showing eventual consistency)
  > in that case rename this 'project' to event-driven-go-demo (can also create an event-driven-py-demo, etc..)
  > might be a better angle (+ add more services // warehouse service // dispatch/delivery service // mark order delivered, etc)
  > database access services? (-> gRPC)
    > requesting coalescing

discussion on protobuf schema of events
  > recently discovered: https://www.asyncapi.com/ (todo: have a look into it. potentially write schemas with it. find way to integrate protobuf as it appears json-oriented)

## Deployment

how to run project
  > write docker compose file (add kafka and kafka ui)

tracing method?
  > some sort of request ID/transaction ID that spans all saga related events (though not needed - tracing could be performed using just order ID)
  > https://medium.com/dzerolabs/observability-journey-understanding-logs-events-traces-and-spans-836524d63172

## License

TODO

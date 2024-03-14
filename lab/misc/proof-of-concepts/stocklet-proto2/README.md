# Stocklet

An event-driven microservices-based distributed e-commerce application written in Go.

## 📘 About

This project was originally made to experiment with using event-driven architecture.

But I hope it can serve as a beneficial demonstration of using the architecture and exemplify the implementation of some other microservice patterns.

Any ideas, contributions or suggestions to better conform with general and evolving industry practices are very welcome and will be greatly appreciated, as I'd like for this project to be somewhat of a reflection of a production-ready enterprise application.

## 📝 Features

* Monorepository layout
* Microservice architecture
* Event-driven architecture
* Schema-driven development
* Interfacing with services using gRPC
* User-facing RESTful HTTP APIs with gRPC-Gateway
* Distributed tracing with OpenTelemetry
* Transactional outbox pattern with Debezium
* API gateway pattern using Envoy
* Choreography-based sagas

TODO: additional features
* ... Idempotent consumers?
* ... domain driven design (not currently true?)

## 🗃️ Architecture

### 🔎 Overview

TODO: diagram

### 🧰 Technical Stack

**Libraries, Frameworks and Tools**

todo: update later

* API Tooling
  * [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc)
  * [github.com/grpc-ecosystem/grpc-gateway/v2](https://pkg.go.dev/github.com/grpc-ecosystem/grpc-gateway/v2)

* Client Libraries
  * [go.opentelemetry.io/otel](https://pkg.go.dev/go.opentelemetry.io/otel)
  * [github.com/twmb/franz-go](https://pkg.go.dev/github.com/twmb/franz-go)
  * [github.com/jackc/pgx/v5](https://pkg.go.dev/github.com/jackc/pgx/v5)

* Protobuf Libraries
  * [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf)
  * [github.com/bufbuild/protovalidate-go](https://pkg.go.dev/github.com/bufbuild/protovalidate-go)
  * [github.com/mennanov/fmutils](https://pkg.go.dev/github.com/mennanov/fmutils)

* Tools
  * [github.com/bufbuild/buf/cmd/buf](https://buf.build/docs/installation)
  * [github.com/golang-migrate/migrate/v4](https://pkg.go.dev/github.com/golang-migrate/migrate/v4#section-readme)

* Miscellaneous
  * [github.com/rs/zerolog](https://pkg.go.dev/github.com/rs/zerolog)
  * [github.com/lestrrat-go/jwx/v2](https://pkg.go.dev/github.com/lestrrat-go/jwx/v2)
  * [github.com/doug-martin/goqu/v9](https://pkg.go.dev/github.com/doug-martin/goqu/v9)

**Infrastructure**

TODO: update later

* Message Brokers
  * [Kafka](https://hub.docker.com/r/bitnami/kafka)
  * [NATS](https://hub.docker.com/_/nats) *(planned support)*
* Databases
  * [PostgreSQL](https://hub.docker.com/_/postgres)
  * [MongoDB](https://hub.docker.com/_/mongo) *(planned support)*
* Miscellaneous
  * [OpenTelemetry](https://opentelemetry.io/)
  * [Envoy](https://www.envoyproxy.io/)
  * [Debezium Connect](https://hub.docker.com/r/debezium/connect)
  * [Debezium Server](https://hub.docker.com/r/debezium/server) *(planned usage)*
* Provisioning and Deployment
  * [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
  * [Kubernetes](https://kubernetes.io/) *(planned usage)*

### 🧩 Services

| Name | gRPC Server | gRPC Gateway | Produces Events | Consumes Events |
| --- | --- | --- | --- | --- |
| [auth](/internal/svc/auth/) | ✔️ | ✔️ | ❌ | ✔️ |
| [order](/internal/svc/order/) | ✔️ | ✔️ | ✔️ | ✔️ |
| TODO | ✔️ | ✔️ | ✔️ | ✔️ |

todo: update service list

Each service is prepared by a [``service-init``](/cmd/service-init/) container, which is responsible for performing any database migrations and configuring the Debezium outbox connectors for that service.

### 📇 Events

Events are serialised in [protocol buffer](https://protobuf.dev/) format. The current event schemas can be located with the following path pattern [``/schema/protobufs/stocklet/<serviceName>/v1/events.proto``](/schema/protobufs/).

Further documentation on the events can be found in [``/docs/events/README.md``](/docs/events/README.md)

## 💻 Deployment

### Using Docker

todo: write-up on deployment with docker compose

### Using Kubernetes

todo: implement support for deploying with kubernetes

## ✍️ License and Contributing

Contributions are always welcome! Please feel free to open an issue or a pull request if you feel you have any ideas for improvement or further expansion of this demo project.

This project is licensed under the [GNU AGPL v3](/LICENSE).
# Stocklet

An event-driven microservices-based distributed e-commerce example application written in Golang.

## 📘 About

This project was originally built as an experiment with event-driven architecture. But I hope it can future serve as a beneficial demonstration of utilising the architecture and exemplify the implementation of some other miscellaneous microservice patterns.

Any ideas, suggestions or direct contributions to better conform with general and evolving industry practices are welcome and will be greatly appreciated, as I'd like for this project to evolve to be somewhat a reflection of a production-ready enterprise application.

The application should be considered in the experimental stage. As such breaking changes are to be expected between future commits to this repo, in order to ease the development process and allow for clean refactoring of the project.

## 📝 Features

* Monorepository layout
* Microservice architecture
* Event-driven architecture
* Interfacing with services using gRPC
* User-facing RESTful HTTP APIs with gRPC-Gateway
* Distributed tracing with OpenTelemetry
* Transactional outbox pattern with Debezium
* API gateway pattern using Envoy
* Distributed transactions utilising the saga pattern

## 🗃️ Architecture

### 🔎 Overview

TODO: diagram

### 🧰 Technical Stack

#### Libraries, Frameworks and Tools

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

* Tools
  * [github.com/bufbuild/buf/cmd/buf](https://buf.build/docs/installation)
  * [github.com/golang-migrate/migrate/v4](https://pkg.go.dev/github.com/golang-migrate/migrate/v4#section-readme)

* Miscellaneous
  * [github.com/rs/zerolog](https://pkg.go.dev/github.com/rs/zerolog)
  * [github.com/lestrrat-go/jwx/v2](https://pkg.go.dev/github.com/lestrrat-go/jwx/v2)
  * [github.com/doug-martin/goqu/v9](https://pkg.go.dev/github.com/doug-martin/goqu/v9)

#### Infrastructure

* Message Brokers
  * [Kafka](https://hub.docker.com/r/bitnami/kafka)
* Databases
  * [PostgreSQL](https://hub.docker.com/_/postgres)
* Miscellaneous
  * [OpenTelemetry](https://opentelemetry.io/)
  * [Envoy](https://www.envoyproxy.io/)
  * [Debezium Connect](https://hub.docker.com/r/debezium/connect)
* Provisioning and Deployment
  * [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)

### 🧩 Services

| Name | gRPC (w/ Gateway) | Produces Events | Consumes Events |
| :-: | :-: | :-: | :-: |
| [auth](/internal/svc/auth/) | ✔️ | ❌ | ✔️ |
| [order](/internal/svc/order/) | ✔️ | ✔️ | ✔️ |
| [payment](/internal/svc/payment/) | ✔️ | ✔️ | ✔️ |
| [product](/internal/svc/product/) | ✔️ | ✔️ | ✔️ |
| [shipping](/internal/svc/shipping/) | ✔️ | ✔️ | ✔️ |
| [user](/internal/svc/user/) | ✔️ | ✔️ | ❌ |
| [warehouse](/internal/svc/warehouse/) | ✔️ | ✔️ | ✔️ |

Each service is prepared by a [``service-init``](/cmd/service-init/) container; responsible for performing any database migrations and configuring Debezium outbox connectors for that service.

### 📇 Events

The events are schemed and serialised using [protocol buffers](https://protobuf.dev/). They are dispatched using the [transactional outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html) with [Debezium](https://debezium.io/) used as a relay to read and publish events from database outbox tables to the message broker.

Further documentation on the events can be found at [``/docs/events/README.md``](/docs/events/README.md)

## 💻 Deployment

### Using Docker

The application can be deployed using [Docker Compose](https://docs.docker.com/compose/) with the config files located in [``/deploy/docker/``](/deploy/docker/)

``docker compose up -f deploy/docker/compose.yaml -f deploy/docker/compose.override.yaml``

## 🧪 Contributing

As stated above, all forms of feedback and contributions are welcome and appreciated!

Some of the current implemented functionality is quite bare-bones. If I come to revisit this project at a later date, I might look to implement other features such as a notification service (i.e. sending emails on updates to order statuses), ensuring idempotetency in the consumers, a more realistic shipping fulfilment process or a frontend user interface.

Have any ideas for improvements? Looking to implement a feature? Please don't hesistate to open an issue to discuss or a pull request with your enhancements.

## 📓 License

This project is licensed under the [GNU Affero General Public License v3](/LICENSE).

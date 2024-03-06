# Outboxie

Transactional outbox processing tool written in Go.

Importable as a [library](https://pkg.go.dev/github.com/hexolan/outboxie) or usable via the CLI.

## Features

* Push-based outbox support (for certain databases)
* Adaptability. Use as a library and adapt for your implementation.
* todo.... full write-up

### Supported Databases

Supported databases. Build your own connector.

* PostgreSQL
* MongoDB (TODO)

### Supported Messaging Systems

Supported messaging systems. Build your own sink.

* Kafka
* NATS (TODO)
* RabbitMQ (TODO)

### Leader Election Mechanisms

* Standalone (no leader election)
* Kubernetes
* Redis Distributed Locks (TODO)
* Database Locking (TODO)
* Raft (TODO)

## Usage

### 1) Library Usage

TODO

### 2) CLI Usage

TODO

## Contributing

Contributions are very welcome and any help would be greatly appreciated.

## License

This project is licensed and avaliable open-source under the [Apache License v2.0.](/LICENSE)
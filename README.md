# Panels

A forum application created using a microservices architecture.

![Panels Banner](/docs/img-1.png)

## About

In current form the application serves as a proof of concept containing bare-bones functionality.

There are definitely aspects I'd like to polish off, along with some ideas for additional functionality that I'd like to implement in spare time, as such this application should be considered a work-in-progress.

## Services

The application is currently comprised of 7 services. Users are served the React site from the ``frontend`` and make calls to the REST API exposed by the ``gateway-service``. 

![Architecture](/docs/img-2.jpg)

Inter-service communication is handled through RPC calls (utilising [gRPC](https://grpc.io/)) and [event sourcing](https://microservices.io/patterns/data/event-sourcing.html) (utilising [Kafka](https://kafka.apache.org/)). 

| Service | Language | Datastores | Description |
| --- | --- | --- | --- |
| [frontend](/services/frontend) | TypeScript (React) | N/A | Web-based user interface |
| [gateway-service](/services/gateway-service) | Golang | N/A | Exposes a HTTP REST API to allow users to communicate with the application from the ``frontend``. |
| [panel-service](/services/panel-service) | Golang | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to panels |
| [post-service](/services/post-service) | Golang | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to posts |
| [user-service](/services/user-service) | TypeScript (Node) | [MongoDB](https://www.mongodb.com/) | Responsible for operations related to users |
| [auth-service](/services/auth-service) | Python | [PostgreSQL](https://www.postgresql.org/) | Responsible for authenticating users |
| [comment-service](/services/comment-service) | Python | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to comments |

## Deployment

### Methods

#### **Using [Kubernetes](https://kubernetes.io/):**

*TODO: finalise documentation*

When deploying to Kubernetes, you will need to deploy Redis, Postgres and Mongo instances off cluster and adjust your configuration.

#### **Using [Docker](https://www.docker.com/):**

*TODO: finalise documentation*

For deployment using Docker Compose, the default container configuration (exposed as environment variables in [/docker-compose.yaml](/docker-compose.yaml)) can be left as is.

This is presuming that the [/docker-compose.override.yaml](/docker-compose.override.yaml) file, which contains specification for the instances that each service requires, is also being used. However, if solely the [/docker-compose.yaml](/docker-compose.yaml) is being used then the configuration will need to be changed to point to your instances of the container databases.

### Configuration

For an outline on the environment variables that each service requires, view the ``README.md`` files in each service folder. 

| Service | Example |
| --- | --- |
| [frontend](/services/frontend) | [.env.example](/services/frontend/.env.example) |
| [gateway-service](/services/gateway-service) | [.env.example](/services/gateway-service/.env.example) |
| [panel-service](/services/panel-service) | [.env.example](/services/panel-service/.env.example) |
| [post-service](/services/post-service) | [.env.example](/services/post-service/.env.example) |
| [user-service](/services/user-service) | [.env.example](/services/user-service/.env.example) |
| [auth-service](/services/auth-service) | [.env.example](/services/auth-service/.env.example) |
| [comment-service](/services/comment-service) | [.env.example](/services/comment-service/.env.example) |

## License

**Acknowledgments:**

* Logo Icon: [Tabler Icons](https://github.com/tabler/tabler-icons) ([MIT License](https://github.com/tabler/tabler-icons/blob/master/LICENSE)) 
* Logo Font: [Oregano](https://fonts.google.com/specimen/Oregano) ([Open Font License](https://scripts.sil.org/cms/scripts/page.php?site_id=nrsi&id=OFL))

This repository is licensed under the [Apache License v2.0](/LICENSE).
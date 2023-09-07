# Panels Documentation

## Contributing

Contributions are welcome. Please feel free to open issues or make a pull request.

## Deployment

### Using [Kubernetes](https://kubernetes.io/):

*TODO: finalise documentation*

When deploying to Kubernetes, you will need to deploy Redis, Postgres and Mongo instances off cluster and adjust your configuration.

### Using [Docker](https://www.docker.com/):

*TODO: finalise documentation*

For deployment using Docker Compose, the default container configuration (exposed as environment variables in [/docker-compose.yaml](/docker-compose.yaml)) can be left as is.

This is presuming that the [/docker-compose.override.yaml](/docker-compose.override.yaml) file, which contains specification for the instances that each service requires, is also being used. However, if solely the [/docker-compose.yaml](/docker-compose.yaml) is being used then the configuration will need to be changed to point to your instances of the container databases.

## Configuration

For an outline on the environment variables that each service requires, view the documentation for the individual services (located in the ``README.md`` files of each service folder). 

| Service | Example |
| --- | --- |
| [frontend](/services/frontend) | [.env.example](/services/frontend/.env.example) |
| [gateway-service](/services/gateway-service) | [.env.example](/services/gateway-service/.env.example) |
| [panel-service](/services/panel-service) | [.env.example](/services/panel-service/.env.example) |
| [post-service](/services/post-service) | [.env.example](/services/post-service/.env.example) |
| [user-service](/services/user-service) | [.env.example](/services/user-service/.env.example) |
| [auth-service](/services/auth-service) | [.env.example](/services/auth-service/.env.example) |
| [comment-service](/services/comment-service) | [.env.example](/services/comment-service/.env.example) |
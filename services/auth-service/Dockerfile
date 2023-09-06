FROM python:3.9-alpine

RUN python -m pip install poetry==1.6.1
RUN poetry config virtualenvs.in-project true

WORKDIR /app

# Install the dependencies
COPY pyproject.toml poetry.lock /app/
RUN poetry install --no-cache --no-root

# Install the package
COPY . .
RUN poetry install --no-cache --only-root

# For Docker healthchecking
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.19/grpc_health_probe-linux-amd64 && chmod +x /bin/grpc_health_probe
HEALTHCHECK CMD ["/bin/grpc_health_probe", "-addr=127.0.0.1:9090"]

EXPOSE 9090
CMD ["poetry", "run", "python", "-m", "auth_service.main"]
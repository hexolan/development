FROM golang:1.21 AS build
WORKDIR /app

# Install required modules
COPY go.mod go.sum ./
RUN go mod download

# Build the service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/service1 ./cmd/service1

# Runtime Environment
FROM gcr.io/distroless/static-debian12
COPY --from=build /bin/service1 .

# Expose ports
EXPOSE 90
EXPOSE 9090

CMD ["./service1"]
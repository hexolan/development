FROM golang:1.21 AS build
WORKDIR /app

# Install required modules
COPY go.mod go.sum ./
RUN go mod download

# Build the service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/service ./cmd/service

# Runtime Environment
FROM gcr.io/distroless/static-debian12
COPY --from=build /bin/service .

EXPOSE 90 9090

CMD ["./service"]
FROM golang:1.21 AS build

WORKDIR /go/src/github.com/hexolan/outboxie

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /bin/outboxie ./cmd/outboxie

FROM gcr.io/distroless/static-debian12
COPY --from=build /bin/outboxie /

ENTRYPOINT ["/outboxie"]
CMD ["--help"]
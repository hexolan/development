FROM rust:1.77 AS builder
WORKDIR /app
COPY . /app
RUN cargo build --release

# Runtime environment (for minifying image size)
FROM gcr.io/distroless/cc-debian12
COPY --from=builder /app/target/release/url_shortener /
CMD ["./url_shortener"]
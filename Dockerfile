# Build stage
FROM golang:1.18-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN apk add curl


# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migrations ./migrations

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
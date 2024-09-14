FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -o billing_service ./cmd/main/

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/billing_service /app/billing_service
COPY --from=builder /app/configs /app/configs

ENTRYPOINT ["/app/billing_service"]
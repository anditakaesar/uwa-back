FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /uwa-back

FROM alpine:3.21

WORKDIR /root/

COPY --from=builder /uwa-back .
COPY --from=builder /app/migrations ./migrations/
COPY --from=builder /app/static ./static

CMD ["./uwa-back"]
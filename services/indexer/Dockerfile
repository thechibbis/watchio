FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY shared/ /app/shared/
COPY services/indexer services/indexer

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/indexer ./services/indexer/cmd/indexer/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/indexer /app/bin/indexer
EXPOSE 50053
ENTRYPOINT ["/app/bin/indexer"]

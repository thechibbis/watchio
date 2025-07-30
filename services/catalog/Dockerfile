FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY shared/ /app/shared/
COPY services/catalog services/catalog

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/catalog ./services/catalog/cmd/catalog/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/catalog /app/bin/catalog
EXPOSE 50052
ENTRYPOINT ["/app/bin/catalog"]

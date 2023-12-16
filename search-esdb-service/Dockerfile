# Build stage
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /tmp/search-esdb

COPY go.mod .
COPY go.sum .

RUN go mod download &&\
    go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
FROM scratch

WORKDIR /app/search-esdb

COPY --from=builder /tmp/search-esdb/main .
COPY --from=builder /tmp/search-esdb/config.yaml .
COPY --from=builder /tmp/search-esdb/data ./data

CMD ["./main"]



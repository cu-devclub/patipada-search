# Build stage
FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /tmp/search-esdb

COPY go.mod .
COPY go.sum .

RUN go mod download &&\
    go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
FROM alpine 

WORKDIR /app/search-esdb

RUN apk add --no-cache ca-certificates tzdata &&\
    apk add --no-cache go

COPY --from=builder /tmp/search-esdb/vendor ./vendor
COPY --from=builder /tmp/search-esdb/main .
COPY --from=builder /tmp/search-esdb/configs ./configs
COPY --from=builder /tmp/search-esdb/data ./data

CMD ["./main"]




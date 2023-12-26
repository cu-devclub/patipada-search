# Build stage
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /tmp/auth

COPY go.mod .
COPY go.sum .

RUN go mod download &&\
    go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
FROM scratch

WORKDIR /app/auth

COPY --from=builder /tmp/auth/main .
COPY --from=builder /tmp/auth/app.env .

CMD ["./main"]



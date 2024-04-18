# FROM golang:1.21 AS builder

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o authApp

FROM scratch

WORKDIR /app
# COPY --from=builder /app/authApp /app
COPY authApp /app
COPY app.env /app

CMD ["/app/authApp"]

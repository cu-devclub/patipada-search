FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY searchApp /app
COPY app.env /app

CMD ["/app/searchApp"]
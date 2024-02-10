FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY searchApp /app
COPY app.env /app
COPY /datasource /app/datasource

CMD ["/app/searchApp"]
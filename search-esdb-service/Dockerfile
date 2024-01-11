FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY searchApp /app
COPY app.env /app
COPY /data /app/data

CMD ["/app/searchApp"]
FROM golang:alpine3.15

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main

CMD ["/app/main"]
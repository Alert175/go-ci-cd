FROM golang:1.19.0-alpine3.16

RUN mkdir /app
RUN mkdir /app/logs

ADD . /app

WORKDIR /app

RUN go build main.go

CMD ["/app/main"]
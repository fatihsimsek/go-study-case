FROM golang:alpine3.14
RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o main ./cmd

CMD ["/app/main"]
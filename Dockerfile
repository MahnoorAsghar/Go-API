FROM golang:1.18

RUN mkdir -p /go-api

WORKDIR /go-api

ADD . /go-api

ENV NAME=mudo

RUN go build ./go-api.go

CMD ["./go-api"]

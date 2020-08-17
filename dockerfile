FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /chat-dg
RUN mkdir -p /chat-dg/proto

WORKDIR /chat-dg

COPY ./proto/service.pb.go /chat-dg/proto
COPY ./main.go /chat-dg

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o chat-dg .

CMD ./chat-dg
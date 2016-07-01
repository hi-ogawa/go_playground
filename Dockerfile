FROM golang:1.6

RUN go get github.com/smartystreets/goconvey

RUN mkdir -p /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

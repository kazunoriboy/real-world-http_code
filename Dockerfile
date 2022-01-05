FROM golang:1.18-rc
RUN mkdir /go/src/work
RUN apt-get update && apt-get install -y vim
WORKDIR /go/src/work
ADD . /go/src/work

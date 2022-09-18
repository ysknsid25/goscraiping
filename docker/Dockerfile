FROM golang:1.19.1-alpine
RUN apk add --update &&  apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
RUN git clone https://github.com/ex-n-soldiers/sample-site.git 
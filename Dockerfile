FROM golang:1.18.2-alpine3.16

RUN apk update && apk upgrade && apk add --update --upgrade bash curl make
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
  mkdir /gobin && \
  mv migrate /gobin && \
  export PATH=$PATH:/gobin

ENV PATH="$PATH:/bin/bash:/gobin"

WORKDIR /go/app
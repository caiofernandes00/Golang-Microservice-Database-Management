FROM golang:1.18.2-alpine3.16

RUN apk add --update --upgrade bash curl make gcc g++
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
  mkdir /gobin && \
  mv migrate /gobin && \
  export PATH=$PATH:/gobin

ENV PATH="$PATH:/bin/bash:/gobin"

# RUN go mod download

WORKDIR /go/app
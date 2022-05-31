FROM golang:1.18.2-alpine3.16

RUN apk add --update --upgrade bash curl make gcc g++

ENV PATH="$PATH:/bin/bash:/gobin"

WORKDIR /go/Golang-Microservice-Database-Management
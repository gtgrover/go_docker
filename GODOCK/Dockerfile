ARG BASE_IMAGE=golang
ARG BASE_TAG=1.12.0-alpine3.9
FROM ${BASE_IMAGE}:${BASE_TAG} as base

LABEL maintainer="GT Grover gtgrover@gmail.com"

RUN apk update --no-cache && apk upgrade --no-cache
RUN apk add --no-cache git
RUN go get github.com/go-redis/redis

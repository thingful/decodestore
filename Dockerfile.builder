FROM golang:1.10-alpine

RUN apk add --update --no-cache \
    git \
    protobuf \
    protobuf-dev && \
  rm -rf /var/cache/apk/* && \
  go get github.com/twitchtv/retool
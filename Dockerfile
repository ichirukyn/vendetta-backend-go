ARG GO_VERSION=1.22.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app/main.exe ./cmd/web

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add chromium

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api .

EXPOSE 5002
EXPOSE 5003
ENTRYPOINT ["/api/app/main.exe"]

FROM golang:1.18.0

LABEL maintaner="adithya"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
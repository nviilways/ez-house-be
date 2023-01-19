FROM golang:1.18.0

LABEL maintaner="adithya"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . ./

RUN go mod download

COPY ./db/house_booking_db.sql ./docker-entrypoint-initdb.d/house_booking_db

ENV GIN_MODE=release

RUN go build -o /build

EXPOSE 8080

CMD [ "/build" ]
FROM golang:1.14 as builder

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./build/tasker

FROM ubuntu:18.04

RUN apt update && apt -y upgrade

WORKDIR /app
COPY --from=builder /app/build/tasker tasker
COPY secrets/environment.tasker.production .env

CMD ["./tasker"]
FROM golang:1.15 as builder

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./build/tasker

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=builder /app/build/tasker tasker
COPY secrets/environment.tasker.production .env

CMD ["./tasker"]
FROM golang:1.15-alpine AS build
WORKDIR /build

RUN apk add --no-cache git build-base

COPY . .
RUN go mod tidy
RUN go build -o bin/lockerr main.go

FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /build/bin .

ENV BOT_TOKEN=""

CMD ./lockerr
FROM golang:1.15-alpine AS build
ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /build

RUN apk add --no-cache git build-base

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# Build the binary without debug information.
RUN go build -o bin/lockerr -ldflags '-w -s' cmd/lockerr/main.go

FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /build/bin .

ENV BOT_TOKEN=""

CMD ./lockerr
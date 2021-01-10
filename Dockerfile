FROM golang:1.15-alpine AS build
WORKDIR /build

RUN apk add --no-cache git build-base upx

COPY . .
RUN go mod tidy
# Build the binary without debug information.
RUN go build -o bin/lockerr -ldflags '-w -s' main.go
# Compress the binary using upx.
RUN upx --lzma bin/lockerr

FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /build/bin .

ENV BOT_TOKEN=""

CMD ./lockerr
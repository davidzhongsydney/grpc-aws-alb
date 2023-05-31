FROM golang:1.19-buster AS builder

RUN mkdir -p /src
COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:buster-slim

RUN mkdir -p /app

COPY --from=builder /src/bin/protobuf /app/

WORKDIR /app

EXPOSE 9000
EXPOSE 80

CMD ["./protobuf"]
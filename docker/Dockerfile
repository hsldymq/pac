FROM golang:1.12-alpine3.10 as temp

COPY go.mod /build/go.mod
COPY go.sum /build/go.sum
COPY main.go /build/main.go
COPY vendor /build/vendor

RUN cd /build && GOFLAGS=-mod=vendor go build -o server

FROM alpine:3.10

COPY --from=temp /build/server /var/www/server
RUN mkdir -p /var/www/pac

ENV LISTEN_PORT 10801

WORKDIR /var/www

CMD ["/var/www/server"]
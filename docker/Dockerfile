FROM golang:1.12-alpine as temp

COPY go.mod /build/go.mod
COPY go.sum /build/go.sum
COPY main.go /build/main.go
COPY vendor /build/vendor

RUN cd /build && GOFLAGS=-mod=vendor go build -o server

FROM golang:1.12-alpine

COPY --from=temp /build/server /var/www/server
RUN mkdir -p /var/www/pac

ENV LISTEN_PORT 10801

WORKDIR /var/www

CMD ["/var/www/server"]
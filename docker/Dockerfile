FROM golang:1.12-alpine3.10 as temp

COPY main.go /build/main.go

RUN cd /build && go build -o server

FROM alpine:3.10

COPY --from=temp /build/server /var/www/server
RUN mkdir -p /var/www/pac

ENV LISTEN_PORT 10801

WORKDIR /var/www

CMD ["/var/www/server"]
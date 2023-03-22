FROM alpine:latest
MAINTAINER kevin <kevinzhang.cq@gmail.com>

WORKDIR /web-server
COPY cmd/web-server /web-server/web-server

EXPOSE 8080

ENTRYPOINT [ "/web-server/web-server" ]

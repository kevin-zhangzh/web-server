FROM alpine:latest
MAINTAINER kevin <kevinzhang.cq@gmail.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /web-server
COPY cmd/web-server /web-server/web-server

EXPOSE 8080

ENTRYPOINT [ "/web-server/web-server" ]

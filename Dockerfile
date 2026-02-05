FROM golang:1.25 AS Builder

COPY src /go
COPY Makefile /go

run make build

FROM scratch

COPY src/www /www

COPY --from=builder /go/src/bin/httpd /httpd

EXPOSE 8888

ENTRYPOINT [ "/httpd" ]
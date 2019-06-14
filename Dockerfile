# Multistage image keeps image size very small 
FROM golang:1.12 AS builder
COPY . /app
WORKDIR /app
RUN go get -d
RUN go get -u github.com/gobuffalo/packr/packr
RUN packr build

FROM alpine:latest
WORKDIR /app
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /app/motd ./
CMD ["./motd"]
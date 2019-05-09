FROM golang:1.12-alpine3.9 as builder

WORKDIR /sample-web-golang
COPY . .

RUN apk add -u git && \
    go build -o /tmp/sample-web main.go

FROM alpine:3.9

COPY --from=builder /tmp/sample-web /usr/bin
WORKDIR /app

EXPOSE 8080

CMD ["sample-web"]

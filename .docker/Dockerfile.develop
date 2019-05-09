FROM golang:1.12-alpine

WORKDIR /sample-web-golang
COPY . .

RUN apk add -u git neovim bash && \
    go get -u github.com/pilu/fresh

CMD ["fresh"]

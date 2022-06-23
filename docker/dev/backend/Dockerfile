FROM golang:1.16-alpine

RUN apk add --no-cache git

RUN go get -u github.com/go-sql-driver/mysql \
  github.com/gin-gonic/gin \
  github.com/gin-contrib/cors

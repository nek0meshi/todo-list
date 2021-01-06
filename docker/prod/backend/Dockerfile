# build-stage

FROM golang:1-alpine as builder

RUN apk add --no-cache git

RUN go get -u github.com/go-sql-driver/mysql \
  github.com/gin-gonic/gin \
  github.com/gin-contrib/cors

WORKDIR /go/src

COPY backend/todo /go/src/todo

RUN go build -o /go/bin/todo todo


# deploy-stage

FROM alpine

COPY --from=builder /go/bin/todo /app

ENTRYPOINT ["/app"]

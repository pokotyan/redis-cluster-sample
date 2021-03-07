FROM golang:1.16.0-alpine3.13 AS builder
WORKDIR $GOPATH/src/redis-cluster-sample
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder $GOPATH/src/redis-cluster-sample/app .
CMD ["./app"]
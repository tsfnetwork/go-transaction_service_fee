# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-transaction_service_fee
RUN cd /go-transaction_service_fee && make gtsf

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-transaction_service_fee/build/bin/gtsf /usr/local/bin/

EXPOSE 4949 4949 59997 59997/udp
ENTRYPOINT ["gtsf"]

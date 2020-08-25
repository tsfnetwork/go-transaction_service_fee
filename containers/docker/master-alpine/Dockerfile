FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/tsfnetwork/go-transaction_service_fee && \
  (cd go-transaction_service_fee && make gtsf) && \
  cp go-transaction_service_fee/build/bin/gtsf /gtsf && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-transaction_service_fee && rm -rf /var/cache/apk/*

EXPOSE 4949
EXPOSE 59997

ENTRYPOINT ["/gtsf"]

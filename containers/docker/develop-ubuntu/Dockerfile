FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/tsfnetwork/go-transaction_service_fee && \
  (cd go-transaction_service_fee && make gtsf) && \
  cp go-transaction_service_fee/build/bin/gtsf /gtsf && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-transaction_service_fee

EXPOSE 4949
EXPOSE 59997

ENTRYPOINT ["/gtsf"]

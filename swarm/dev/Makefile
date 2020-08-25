.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/tsfnetwork/go-transaction_service_fee/cmd/swarm
	go build -o bin/geth     github.com/tsfnetwork/go-transaction_service_fee/cmd/geth
	go build -o bin/bootnode github.com/tsfnetwork/go-transaction_service_fee/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/tsfnetwork/go-transaction_service_fee/swarm/...

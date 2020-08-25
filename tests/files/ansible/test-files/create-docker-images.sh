#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="transaction_service_fee
/cppjit-testrunner" docker-cppjit
docker build --tag="transaction_service_fee
/python-testrunner" docker-python
docker build --tag="transaction_service_fee
/go-testrunner" docker-go

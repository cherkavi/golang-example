#!/bin/sh

GOOS=linux GOARCH=amd64 go build
docker build -t smartparkregistry.stritzke.me/smartpark/service:latest .
docker push smartparkregistry.stritzke.me/smartpark/service:latest

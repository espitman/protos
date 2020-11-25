#!/bin/bash 

protoc --go_out=plugins=grpc:. acm/acm.proto \
&& protoc-go-inject-tag -input=./acm/acm.pb.go \
&& git add . && git commit -m 'acm update' && git push
#!/bin/bash 

protoc --go_out=plugins=grpc:. amenity/amenity.proto \
&& protoc-go-inject-tag -input=./amenity/amenity.pb.go \
&& git add . && git commit -m 'amenity update' && git push
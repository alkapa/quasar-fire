#!/usr/bin/env bash

for p in $( ls ./proto/ ); do
    protoc \
		-I ./proto/ \
		-I $GOPATH/src/github.com/alkapa/ \
		-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/ \
		--plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway \
		--go_out=$GOPATH/src \
		--go-grpc_out=require_unimplemented_servers=false:$GOPATH/src \
    --grpc-gateway_out=logtostderr=true:$GOPATH/src \
		--openapiv2_out=:swagger-ui \
		--openapiv2_opt=logtostderr=true \
		proto/$p
done

#!/bin/bash

sudo PATH=$PATH:$GOPATH/bin/ protoc --go_out=plugins=grpc:. common/model/article.proto

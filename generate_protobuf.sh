#!/usr/bin/env zsh

protoc api/chat/chat.proto --go_out=plugins=grpc:.

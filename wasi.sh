#!/bin/sh

tinygo \
	build \
	-o ./imports2json.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./cmd/imports2json/main.go

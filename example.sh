#!/bin/sh

cat parser.go |
	wazero run ./imports2json.wasm |
	jq -c '.[]'

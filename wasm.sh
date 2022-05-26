#!/bin/bash

# build for linux
GOOS=js GOARCH=wasm go build -o ./build/sail.wasm iatearock.com/has


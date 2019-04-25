#!/bin/bash

for x in **/*.proto; do protoc --go_out=paths=source_relative:. --plugin=grpc $x; done
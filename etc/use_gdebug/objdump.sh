#!/bin/bash
set -eu
go build -gcflags "-N -l" ./helloecho.go
go tool objdump -s "counting" ./helloecho > ./objdump_couting.log


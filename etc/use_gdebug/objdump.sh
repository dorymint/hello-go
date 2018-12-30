#!/bin/bash
set -eu

# cds(change to script directory)
cd "$(dirname "$(readlink -f "${0}")")"

go build -gcflags "-N -l" ./helloecho.go
go tool objdump -s "counting" ./helloecho > ./objdump_couting.log


#!/bin/bash
rm -rf cmd/
rm -rf gen/
rm -f simpleapi
goa gen simpleapi/design
goa example simpleapi/design
go build ./cmd/simpleapi

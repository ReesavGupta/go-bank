#!/bin/bash
set -e

go build -o bin/gobank
./bin/gobank

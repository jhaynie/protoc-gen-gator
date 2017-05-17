#!/usr/bin/env bash

set -euo pipefail

protoc -I. \
--go_out=:. \
*.proto

#!/usr/bin/env bash

GOOSARCH="${GOOS}_${GOARCH}"
GOOS="linux"

$cmd docker build --tag generate:$GOOS $GOOS
$cmd docker run --interactive --tty --volume $(cd -- "$(dirname -- "$0")" && /bin/pwd):/build generate:$GOOS

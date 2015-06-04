#!/bin/bash

set -ex

cd "$(dirname "$(readlink -f "$0")")"

[ -d .coverage ] || mkdir .coverage

# Find subpackages:
for pkgpath in $(find . -mindepth 2 -type f -name '*.go' -print0 | xargs -0 -n 1 dirname | sort -u)
do
    coverdata=".coverage/$(echo $pkgpath | sed 's,^\./,,; s,/,_,g').data"
    go test -covermode=count -coverprofile="$coverdata" "$pkgpath"
    [ -f "$coverdata" ] && go tool cover -html="$coverdata"
done

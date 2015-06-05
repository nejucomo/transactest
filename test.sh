#!/bin/bash

#set -x
set -eu

# Use https://github.com/matm/gocov-html because `go tool cover` can't
# seem to generate an inert html file without trying to invoke a browser
# process. (I also tried to find a ticket tracker to suggest this breaks
# most automated report generation, but could not find one.) >:S

cd "$(dirname "$(readlink -f "$0")")"


# Go package / local path convolution is cumbersome. I hardcode this here,
# and assume the git checkout is on your $GOPATH.
PKGNAME='github.com/nejucomo/transactest'
SRCROOT=''
for d in "$(echo "$GOPATH" | tr ':' '\n')"
do
    t="$d/src"
    if [ -d "$t/$PKGNAME" ]
    then
        SRCROOT="$t"
        break
    fi
done

if [ -z "$SRCROOT" ]
then
    echo "Could not find go package $PKGNAME on \$GOPATH=$GOPATH"
    exit 1
fi

[ -d .coverage ] || mkdir -v .coverage

# Find subpackages:
for subpkg in $(cd "$SRCROOT"; find -L "$PKGNAME" -mindepth 2 -type f -name '*.go' -print0 | xargs -0 -n 1 dirname | sort -u)
do
    reportname="$(echo "$subpkg" | sed 's,/,_,g')"
    reportjson=".coverage/${reportname}.json"
    reporthtml=".coverage/${reportname}.html"

    gocov test "$subpkg" > "$reportjson"

    # Work-around a case gocov-html does not handle...
    if [ -f "$reportjson" ] && ! grep -q '^{"Packages":null}$' "$reportjson"
    then
        echo "  Generating coverage report: $reporthtml"
        gocov-html "$reportjson" > "$reporthtml"
        echo
    fi
done

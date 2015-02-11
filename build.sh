#!/bin/bash

# Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
# Use of this source code is governed by a GPL-style
# license that can be found in the LICENSE file.
#
# Build script for to.t34.me
# -v - verbose mode
# -f - force mode

program="go.t34.me"
gobin="`which go`"
gitbin="`which git`"
repo="src/github.com/z0rr0/go.t34.me"

if [ -z "$GOPATH" ]; then
	echo "ERROR: set $GOPATH env"
	exit 1
fi
if [ ! -x "$gobin" ]; then
	echo "ERROR: can't find 'go' binary file"
	exit 2
fi
if [ ! -x "$gitbin" ]; then
	echo "ERROR: can't find 'git' binary file"
	exit 3
fi

cd ${GOPATH}/${repo}
gittag="`$gitbin tag | sort --version-sort | tail -1`"
gitver="`$gitbin log --oneline | head -1 `"
build="`date --utc +\"%F %T\"` UTC"
version="$gittag git:${gitver:0:7} $build"

options=""
while getopts ":fv" opt; do
    case $opt in
        f)
            options="$options -a"
            ;;
        v)
            options="$options -v"
            echo "$program version: $version"
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            ;;
    esac
done

$gobin install $options -ldflags "-X main.Version \"$version\"" github.com/z0rr0/go.t34.me/main
exit 0


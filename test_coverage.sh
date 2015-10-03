#!/bin/bash
# Copyright (c) 2015, Peter Mrekaj. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE.txt file.

# This script helps gathering and send (to Covaralls) test coverage across
# multiple Golang packages listed in the same directory as this script is located.

go get github.com/axw/gocov/gocov
go get github.com/mattn/goveralls
go get golang.org/x/tools/cmd/cover

echo "mode: set" > cc.out
shopt -s nullglob
for dir in $(find . -type d); do
    set -- $dir/*.go
    if [ "$#" -gt 0 ]; then
        go test -coverprofile=profile.out $dir
        if [ $? -eq 0 ]; then
            if [ -f profile.out ]; then
                grep -v "mode: set" profile.out >> cc.out
            fi
        else
            rm -f cc.out profile.out
            exit 1
        fi  
    fi
done

$HOME/gopath/bin/goveralls -service=travis-ci -coverprofile=cc.out

rm -f cc.out profile.out

exit 0

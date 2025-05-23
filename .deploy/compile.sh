#!/bin/bash

source ~/.profile

cd "$HOME"/compile/datecalc-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
pwd
echo "------------------------------------"

APPRELEASEVERSION=$(git rev-list -1 HEAD)
export APPRELEASEVERSION
echo "REV: $APPRELEASEVERSION"

echo ""
echo "compiling: go build datecalc.go -ldflags -X main.REV=$APPRELEASEVERSION"
go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/datecalc || {
    echo "Status: $?"
    exit 4
}

echo ""
echo "compiling: GOOS=aix GOARCH=ppc64 go build datecalc.go -ldflags -X main.REV=$APPRELEASEVERSION"
GOOS=aix GOARCH=ppc64 go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/datecalc.aix || {
    echo "Status: $?"
    exit 4
}

echo ""
echo "compiling: GOOS=windows GOARCH=amd64 go build datecalc.go -ldflags -X main.REV=$APPRELEASEVERSION"
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/datecalc.win64 || {
    echo "Status: $?"
    exit 4
}

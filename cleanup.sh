#!/bin/bash -e

tag=$1

for m in $(find ./ -name 'go.mod'); do
  d=$(dirname $m);
  pushd $d;
  grep -q github.com/geiqin/go-micro go.mod && go get github.com/geiqin/go-micro@$tag
  grep -q github.com/geiqin/micro go.mod && go get github.com/geiqin/micro@$tag
  go mod tidy
  popd;
done

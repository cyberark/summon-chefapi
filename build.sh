#!/bin/bash

rm -rf pkg

docker build -t summon/build .

projectpath="/goroot/src/github.com/conjurinc/summon-chefapi"
buildcmd='GOPATH="/goroot/src/github.com/conjurinc/summon-chefapi/Godeps/_workspace:/gopath" GOX_OS="darwin linux windows" GOX_ARCH="amd64" gox -verbose -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"'

docker run --rm \
-v "$(pwd)":"${projectpath}" \
-w "${projectpath}" \
summon/build \
bash -c "${buildcmd}"
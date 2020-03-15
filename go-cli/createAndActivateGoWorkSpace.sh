#! /bin/bash

# current directory is workspace
WORKSPACE=$(pwd)
cd "$WORKSPACE"
mkdir -p src bin pkg

export GOPATH=$WORKSPACE
export GOBIN=$WORKSPACE/bin

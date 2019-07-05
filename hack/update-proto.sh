#!/bin/bash

# Only use this when testing a new API before commiting the API repository. 
# Use dep ensure as the primary way to keep the GRPC code in sync.

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."

cd $root

rsync -a ../apis/grpc/pki/ vendor/github.com/linuxuser586/apis/grpc/pki/

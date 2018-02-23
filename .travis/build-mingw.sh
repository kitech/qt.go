#!/bin/bash

# stop as soon as one of steps will fail
set -e -o pipefail

docker build -t archmingwqt -f Dockerfile.mingw .

ls
docker run -v $HOME/.ccache:/root/.ccache archmingwqt $WINARCH


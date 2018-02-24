#!/bin/bash

# stop as soon as one of steps will fail
set -e -o pipefail

docker build -t archmingwqt -f Dockerfile.ubuntu16 .

ls
docker run -v $HOME/.ccache:/root/.ccache archmingwqt $GOVER


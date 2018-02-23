#!/bin/bash

# Fail out on error
set -e -o pipefail

# before_install:
sudo apt-get -qq update
sudo apt-get install -y libffi-dev

# before_script:
go get -v github.com/gonuts/dl
go get -v github.com/emirpasic/gods/lists/arraylist
go get -v github.com/thoas/go-funk
go get -v github.com/kitech/goplusplus

# script:
pwd
make qtrt- bases qmls extras tools


#!/usr/bin/env bash

set -euxo pipefail

sudo apt -yq update
sudo apt -yq upgrade
sudo apt -yq install build-essential git htop

# install chrome
# install Go
# add go to path in .profile

#!/bin/bash
# First, install the go file for Linux at  https://go.dev/dl/
sudo apt install -y golang

# Then add the following to your .bashrc
export GOROOT=/usr/lib/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Reload your .bashrc
source ~/.bashrc

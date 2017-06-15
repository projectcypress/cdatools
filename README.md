cdatools [![Build Status](https://travis-ci.org/projectcypress/cdatools.svg?branch=master)](https://travis-ci.org/projectcypress/cdatools)
================================
# Installation

`go get github.com/projectcypress/cdatools/...`

# Building

You will also need go-bindata which you can install with

`go get -u github.com/jteeuwen/go-bindata/...`

Make sure that $GOPATH/bin is in your path and then run `make`

Or if you are doing active development you can use `make debug` which will read the contents of the file each time it's accessed so you don't need to keep rerunning make while building templates.

#!/bin/sh

docker run -it \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    golang


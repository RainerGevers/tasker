#!/usr/bin/env bash

version=$(cat VERSION)

docker build -t rainerza/tasker:latest -t rainerza/tasker:$version .
docker push rainerza/tasker:latest rainerza/tasker:$version .
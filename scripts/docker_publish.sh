#!/bin/bash
# if the first argument is empty abort
if [ -z "$1" ]; then
    echo "No argument supplied, must supply tag"
    exit 1
fi
# This script is used to publish the docker image to docker hub, eventually move to github actions
docker build -t tinymonitor:latest .
docker tag tinymonitor:latest lbrictson/tinymonitor:$1
docker push lbrictson/tinymonitor:$1
docker build -t tinymonitor:slim-latest -f slim_Dockerfile .
docker tag tinymonitor:slim-latest lbrictson/tinymonitor:$1-slim
docker push lbrictson/tinymonitor:$1-slim
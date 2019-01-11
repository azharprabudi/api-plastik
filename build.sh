#!/bin/bash

docker-compose build
docker rm -f $(docker ps -a -q -f status=exited)
docker rmi -f $(docker images -q -f dangling=true)
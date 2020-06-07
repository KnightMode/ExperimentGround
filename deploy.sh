#!/usr/bin/env bash
cat creds.txt | docker login -u ${CI_REGISTRY_USER} --password-stdin
mkdir -p app && cd app
docker pull shreegb/docker_deploy_test
docker container stop sample-api
docker container rm sample-api
docker rmi shreegb/docker_deploy_test -f
docker container run --name sample-api -p 8080:8080 -d shreegb/docker_deploy_test
#!/usr/bin/env bash
cat creds.txt | docker login -u ${CI_REGISTRY_USER} --password-stdin
mkdir -p app && cd app
docker pull shreegb/docker_deploy_test
docker container run -p 8080:8080 -d shreegb/docker_deploy_test
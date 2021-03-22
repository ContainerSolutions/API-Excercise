#!/bin/sh

# Before running this script, have installed kind, docker, and kubectl.
# This script sets up the local registry used for
# pushing and pulling docker images.
# The workflow reflects the contents of the GitHub Action
# in order for developing on local to resemble the
# environment in portion of the pipeline where the
# GitHub Action lives.
# As implied by the README, the order of scripts run when
# developing on local is:
#   - run ./setup_local.sh
#   - run ./build.sh
#   - run ./push.sh
#   - run ./deploy.sh
# Prior to running this script, you may want to
#   - delete existing kind cluster with the name 'api-cluster'
#   - stop docker processes named 'api-cluster-worker' and
# 'api-cluster-control-plane'
#   - remove docker container with name '/kind-registry'

kind create cluster --config .ci/kind.yaml --name api-cluster

docker run -d --restart=always -p "127.0.0.1:5000:5000" --name kind-registry registry:2

kubectl apply -f .ci/ingress.yaml

kubectl -n ingress-nginx wait --for=condition=ready --timeout=600s pod \
    -l app.kubernetes.io/name=ingress-nginx,app.kubernetes.io/component=controller

docker pull nginx:alpine
docker tag nginx:alpine kind-registry:5000/nginx:alpine
docker push kind-registry:5000/nginx:alpine

echo
echo "Ensure the following lines exist in /etc/hosts"
echo "127.0.0.1 api.awesome"
echo "127.0.0.1 kind-registry"

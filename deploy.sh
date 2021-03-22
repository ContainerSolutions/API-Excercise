#!/bin/bash

# The kubernetes manifests pull the docker images from local registry.
# The resources have all been labelled with 'group=titanic'.

kubectl apply -f manifests/

kubectl wait --for=condition=ready --timeout=300s deployment -l group=titanic

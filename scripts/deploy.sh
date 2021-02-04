#!/bin/sh
echo "Deploying Kubernetes cluster"

kubectl apply -f k8s_manifests/. -R

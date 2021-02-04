#!/bin/bash

kubectl apply -f k8s_manifests/01-ns.yaml
kubectl apply -f k8s_manifests/backend/
kubectl apply -f k8s_manifests/storage/

sleep 10

kubectl -n titanic get pods

#!/bin/bash

kubectl apply -f k8s_manifests/01-ns.yaml
kubectl apply -f k8s_manifests/backend/
kubectl apply -f k8s_manifests/storage/

kubectl -n titanic wait --for=condition=ready --timeout=600s pod -l app=postgres
kubectl -n titanic wait --for=condition=ready --timeout=600s pod -l app=titanic


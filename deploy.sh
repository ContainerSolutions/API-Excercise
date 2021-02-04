#!/bin/bash

kubectl apply -f k8s_manifests/01-ns.yaml
kubectl apply -f k8s_manifests/backend/
kubectl apply -f k8s_manifests/storage/


sleep 30

kubectl -n titanic wait --for=condition=ready --timeout=300s pod -l app=postgres
sleep 30
kubectl -n titanic describe deployment app=postgres
kubectl -n titanic wait --for=condition=ready --timeout=300s pod -l app=titanic


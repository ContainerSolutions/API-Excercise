#!/bin/sh

# This script will build a backend image and push to dockerhub.
# Use `docker login` before running it.

if (( $# < 1 ))
then
  echo "Error: Missing argument <tag>"
  echo "Usage: ./scripts/build.sh v45"
  exit 1
fi

IMAGE_NAME="anonymouzaccount/titanic-backend"
TAG="${1}"

echo "Building image ${IMAGE_NAME}:${TAG}"
docker build -t ${IMAGE_NAME}:${TAG} .
docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:latest

echo "Pushing image ${IMAGE_NAME}:${TAG} to dockerhub registry."
docker push ${IMAGE_NAME}:${TAG}
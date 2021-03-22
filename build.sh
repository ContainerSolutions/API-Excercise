#!/bin/sh

# Build docker images for the two services used.
# It is good practice to delete previous images with the
# same name.

# build application server
docker build -t kind-registry:5000/titanic:1.0.0 .

# build seeded db
docker build -t kind-registry:5000/titanic-postgres:1.0.0 -f Dockerfile.db .

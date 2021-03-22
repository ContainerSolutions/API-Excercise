#!/bin/bash

# Push the two docker images to the
# local registry, which should be running.
# If the local registry is not running, see
# the bash script './setup_local.sh'.

docker image push kind-registry:5000/titanic:1.0.0

docker image push kind-registry:5000/titanic-postgres:1.0.0

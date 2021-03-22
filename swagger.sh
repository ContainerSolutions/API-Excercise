#!/bin/bash

# These commands run the swagger docker image to generate the node.js server
# implementing the API specified in the swagger.yml file

docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
-i local/swagger.yml \
-l nodejs-server \
-o local/swagger/

sudo chown $USER -R ./swagger

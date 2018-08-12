#!/bin/bash

# Variables
image="titanic-app"

echo "Building new image with tag: $TAGNAME"
docker build -t $image .
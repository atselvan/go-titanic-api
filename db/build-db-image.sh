#!/bin/bash

# Variables
image="titanic-db"

echo "Building new image with tag: $TAGNAME"
docker build -t $image .
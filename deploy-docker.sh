#!/bin/bash
#
# Create Isolated Network
docker network create isolated_nw
#
# Deploy DB
#cd db && ./build-db-image.sh && cd ..
#db/run-db-container.sh
#
# Deploy App
cd app && ./build-app-image.sh && cd ..
app/run-app-container.sh
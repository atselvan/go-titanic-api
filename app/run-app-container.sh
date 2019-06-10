#!/bin/bash

# Variables
image="titanic-app"
name="titanic-app"
network="isolated_nw"
port="8000:8000"

running=`docker ps | grep -c $name`
if [ $running -gt 0 ]
then
   echo "Stopping $name instance"
   docker stop $name
fi

existing=`docker ps -a | grep -c $name`
if [ $existing -gt 0 ]
then
   echo "Removing $name container"
   docker rm $name
fi

echo "Running a new instance with name $name"
echo "[INFO] IMAGE   : $image"
echo "[INFO] NAME    : $name"
echo "[INFO] NETWORK : $network"
echo "[INFO] PORT    : $port"

docker run --name $name -d -p $port -e DB_HOST=titanic-db -e DB_NAME=titanic -e DB_USER=postgres -e DB_PASSWORD=password --network $network $image
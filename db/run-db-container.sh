#!/bin/bash

# Variables
image="titanic-db:latest"
name="titanic-db"
network="isolated_nw"
port="5432:5432"

running=`docker container ls | grep -c $name`
if [ $running -gt 0 ]
then
   echo "Stopping $name instance"
   docker stop $name
fi

existing=`docker container ls -a | grep -c $name`
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

docker run --name $name -d -p $port --network $network $image
sleep 10
docker exec -it $name gosu postgres psql titanic -f /tmp/create-table.sql
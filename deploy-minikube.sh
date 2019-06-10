#!/bin/bash
#
# Cleanup deployments if any
kubectl delete -f app/deployment.yml
kubectl delete -f db/deployment.yml
# 
# Create Secret for the DB
kubectl create secret generic titanic-db-password --from-literal=password=password
#
# Create image for the application local to minikube
eval $(minikube docker-env)
cd app && ./build-app-image.sh && cd ..
#
# Deploy DB
kubectl apply -f db/deployment.yml
sleep 15 # wait for database to start
#
# Load data into the DB
kubectl cp db/titanic.csv $(kubectl get pods | grep titanic-db | awk '{print $1}'):/tmp titanic.csv
kubectl cp db/create-table.sql $(kubectl get pods | grep titanic-db | awk '{print $1}'):/tmp/create-table.sql
kubectl exec -it $(kubectl get pods | grep titanic-db | awk '{print $1}') -- gosu postgres psql titanic -f /tmp/create-table.sql
#
# Deploy App
kubectl apply -f app/deployment.yml
sleep 5 # wait for app to start
# Open the app
minikube service titanic-app
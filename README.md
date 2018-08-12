# Go Titanic application

## Running the application

Create a docker network called isolated_nw.

```bash
docker network create isolated_nw
```

Run the script db/build-db-image.sh to create a docker image for the titanic database and then run the container by running the image db/run-db-container.sh (This script also loads the data of the passengers from the titanic.csv file).

Next, run the script app/build-app-image.sh to create a image for the application. This image build and runs a go application that connects to the database and exposes a rest API to get the data of the passengers from the titanic. To run the application run the script app/run-app-container.sh.
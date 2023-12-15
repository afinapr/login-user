#! /bin/bash


export PROJECT_ID=poc-ddb-mlops
export REGION=asia-southeast2
export CONNECTION_NAME=poc-ddb-mlops:asia-southeast2:bri-ct-hc-data


gccloud build submit \
--tag gcr.io/$PROJECT_ID/poll \
--project $PROJECT_ID

gcloud run deploy poll \
--image gcr.io/$PROJECT_ID/poll \
--platform managed \
--region $REGION \
--allow-unauthenticated \ 
--add-cloudsql-instances $CONNECTION_NAME \
--project $PROJECT_ID \
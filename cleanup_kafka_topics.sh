#!/bin/bash

# Define the array of ids
ids=(201 202 203 302 305 351 352 353 354 355 356 357 358 501 503 504 551 552 553 901 902 931 932 941 942 951)

# Loop through all ids and delete Kafka topics
for id in "${ids[@]}"; do
  kubectl delete kafkatopic transport-${id} -n kafka
  echo "Deleted KafkaTopic transport-${id}"
done

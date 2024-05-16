#!/bin/bash

# Define the array of ids
ids=(201 202 203 302 305 351 352 353 354 355 356 357 358 501 503 504 551 552 553 901 902 931 932 941 942 951)

# Loop through all ids and create Kafka topics
for id in "${ids[@]}"; do
  cat << EOF | kubectl create -n kafka -f -
apiVersion: kafka.strimzi.io/v1beta2
kind: KafkaTopic
metadata:
  name: transport-${id}
  labels:
    strimzi.io/cluster: "my-cluster"
spec:
  partitions: 1
  replicas: 1
  #config:
  #  retention.ms: 7200000
  #  segment.bytes: 1073741824
EOF
  echo "Created KafkaTopic transport-${id}"
done

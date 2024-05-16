#!/bin/bash

# Define arrays of types and ids
types=("transport")
ids=(201 202 203 302 305 351 352 353 354 355 356 357 358 501 503 504 551 552 553 901 902 931 932 941 942 951)

# Loop through all types and ids
for type in "${types[@]}"; do
    for id in "${ids[@]}"; do
        # Define environment variables for substitution
        export TYPE=$type
        export ID=$id

        # Create a file name
        filename="${type}-${id}-connector.yaml"

        # Use cat and envsubst to generate the YAML
        cat <<EOF | envsubst > "$filename"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${TYPE}-${ID}-connector
  namespace: kafka
  labels:
    app: go-mqtt-kafka-connector
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-mqtt-kafka-connector
  template:
    metadata:
      labels:
        app: go-mqtt-kafka-connector
    spec:
      containers:
      - name: ${TYPE}-${ID}-connector
        image: noyusu/go-mqtt-kafka-connector:v0.0.8
        args: ["--type=${TYPE}", "--id=${ID}"]
EOF
        echo "Generated ${filename}"

        sleep 1
        # Deploy the generated YAML to Kubernetes
        kubectl create -f "${filename}"
        echo "Deployed ${filename} to Kubernetes"
        sleep 1
    done
done

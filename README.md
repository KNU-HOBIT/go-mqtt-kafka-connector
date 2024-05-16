# Go MQTT to Kafka Connector

This repository contains a Go application that bridges MQTT messages to a Kafka topic. The application subscribes to MQTT topics and publishes the received messages directly to Kafka, ensuring seamless data flow between MQTT brokers and Kafka clusters.

## Features

- **MQTT to Kafka Bridging**: Subscribes to specified MQTT topics and forwards messages to corresponding Kafka topics.
- **Dynamic Configuration**: Supports dynamic creation of Kafka topics and Kubernetes deployments via shell scripts.
- **Lightweight Docker Container**: Utilizes a minimal Docker image for efficient deployment.
- **Kubernetes Integration**: Includes scripts for creating and cleaning up Kubernetes resources.

## Project Structure

- `main.go`: The main Go application code.
- `config.yaml`: Configuration file for MQTT and Kafka settings (ignored in Git).
- `Dockerfile`: Dockerfile for building the application container.
- `create_deployments.sh`: Script to create Kubernetes deployments for different MQTT topics.
- `cleanup_deployments.sh`: Script to clean up Kubernetes deployments.
- `create_kafka_topics.sh`: Script to create Kafka topics dynamically.
- `cleanup_kafka_topics.sh`: Script to clean up Kafka topics.

## Getting Started

To set up and run this project for the first time, follow these steps:

1. **Initialize the Go module**:
   ```sh
   go mod init <module-name>
   ```
   Replace <module-name> with your project's module path or name.
2. **Download the dependencies**:
   ```sh
   go mod download
   ```
3. **Run the project**:
   ```sh
   go run main.go --type=<type> --id=<id>
   ```
   Replace <type> and <id> with the appropriate values. For example:
   ```sh
   go run main.go --type=transport --id=932
   ```


### Docker and Kubernetes
1. **Build the Docker Image**:
   ```sh
   docker build -t yourname/go-mqtt-kafka-connector:v0.x.x .
   ```

2. **Push**:
   ```sh
   docker push yourname/go-mqtt-kafka-connector:v0.x.x
   ```
3. **Deploy to Kubernetes**:
   ```sh
   ./create_deployments.sh
   ```
4. **Clean Up Kubernetes Deployments**:
   ```sh
   ./cleanup_deployments.sh
   ```
5. **Create Kafka Topics**:
   ```sh
   ./create_kafka_topics.sh
   ```
6. **Clean Up Kafka Topics**:
   ```sh
   ./cleanup_kafka_topics.sh
   ```

##  Prerequisites
- Docker
- Kubernetes cluster with kubectl configured
- Kafka cluster managed by Strimzi
- MQTT broker

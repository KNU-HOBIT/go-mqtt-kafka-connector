# (base) scouter@scouter1:~/yousoo/strimzi-kafka-connect/go-connector$ tree -L 1
# .
# ├── cleanup_deployments.sh
# ├── cleanup_kafka_topics.sh
# ├── config.yaml
# ├── create_deployments.sh
# ├── create_kafka_topics.sh
# ├── Dockerfile
# ├── go.mod
# ├── go.sum
# └── main.go

# docker build -t noyusu/go-mqtt-kafka-connector:v0.0.8 .
# docker push noyusu/go-mqtt-kafka-connector:v0.0.8

# Start from a Golang base image
FROM golang:1.22.2

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
COPY config.yaml ./

# Build
RUN GOOS=linux go build -o /main

# Run
ENTRYPOINT ["/main"]
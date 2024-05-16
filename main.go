package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v2"
)

// Config stores configuration data from config.yaml
type MQTTConfig struct {
	Broker   string `yaml:"broker"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}
type KAFKAConfig struct {
	Servers string `yaml:"servers"`
}

type Config struct {
	Mqtt  MQTTConfig  `yaml:"mqtt"`
	Kafka KAFKAConfig `yaml:"kafka"`
}

func main() {
	// Reading command line arguments
	topicType := flag.String("type", "", "Topic type for MQTT and Kafka")
	topicID := flag.String("id", "", "Topic ID for MQTT and Kafka")
	flag.Parse()

	if *topicType == "" || *topicID == "" {
		fmt.Println("Usage Example: $ go run main.go --type=transport --id=932")
		os.Exit(1)
	}

	// Reading configuration from YAML file
	var config Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %s\n", err)
		os.Exit(1)
	}

	//for Debugging
	fmt.Printf("topicType: %s\n", *topicType)
	fmt.Printf("topicID: %s\n", *topicID)
	fmt.Printf("MQTT Broker: %s\n", config.Mqtt.Broker)
	fmt.Printf("MQTT User Name: %s\n", config.Mqtt.UserName)
	fmt.Printf("MQTT Password: %s\n", config.Mqtt.Password)
	fmt.Printf("Kafka Servers: %s\n", config.Kafka.Servers)

	mqttTopic := *topicType + "/" + *topicID
	mqttClientID := mqttTopic + "-client"
	kafkaTopic := *topicType + "-" + *topicID

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  config.Kafka.Servers,
		"acks":               "all",
		"enable.idempotence": "true",
		"compression.type":   "lz4",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	/* for Debugging
	var dotCount int = 0

	// In MQTT message handler..
	dotCount++ // Increment the dot counter each time a message is received
	if dotCount >= 100 {
		fmt.Println() // Print a newline after 100 dots
		dotCount = 0  // Reset the counter
	} else {
		fmt.Print(".") // Print a dot without newline
	}
	*/

	// MQTT message handler
	mqttMessageHandler := func(client mqtt.Client, msg mqtt.Message) {
		// Send message to Kafka
		if err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Value:          msg.Payload(),
		}, nil); err != nil {
			fmt.Printf("Failed to produce to Kafka: %s\n", err)
		}
	}

	opts := mqtt.NewClientOptions().AddBroker(config.Mqtt.Broker).SetClientID(mqttClientID)
	opts.SetUsername(config.Mqtt.UserName)
	opts.SetPassword(config.Mqtt.Password)

	// MQTT connect
	opts.SetDefaultPublishHandler(mqttMessageHandler)
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Printf("Error subscribing to topic %s: %s\n", mqttTopic, token.Error())
		}
	}

	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to MQTT broker: %s\n", token.Error())
		os.Exit(1)
	}
	defer mqttClient.Disconnect(250)

	// Wait for messages
	fmt.Println("Waiting for messages...")
	for {
		time.Sleep(time.Second)
	}
}

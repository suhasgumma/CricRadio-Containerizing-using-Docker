package kafka

import (
	"cricradio-go-svc/logger"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)
import "github.com/segmentio/kafka-go"

var (
	host           = "kafka"
	port           = "9092"
	ControllerConn *kafka.Conn
	brokerUrl      string
)

func init() {
	brokerUrl = "kafka:9092"
	//Dialing broker with host & port
	conn, err := kafka.Dial("tcp", brokerUrl)
	if err != nil {
		logger.Error("Error dialing kafka broker : ", err)
	}
	defer conn.Close()

	// checking if successful connection occured or not by getting controller
	controller, err := conn.Controller()
	if err != nil {
		logger.Error("Error with kafka broker connection : ", err)
	}

	ControllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		logger.Error("Error with kafka controller connection : ", err)
	}

	log.Println(fmt.Sprintf("Kafka Broker Connection is %s:%d is Up!", controller.Host, controller.Port))
}

func CreateTopic(topicName string) {

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topicName,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err := ControllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		logger.Error("Error with kafka topic creation : ", err)
	}

	logger.Info(fmt.Sprintf("Topic Created Successfully - %v", topicName))
}

func DeleteTopic(topicName string) {

	err := ControllerConn.DeleteTopics(topicName)
	if err != nil {
		logger.Error("Topic Deletion Failed : ", err)
	}

	logger.Info(fmt.Sprintf("Topic Deletion Successful - %v", topicName))
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func ProduceComm(msg string, topic string, ctx context.Context) {

	writer := newKafkaWriter(brokerUrl, topic)
	defer writer.Close()

	logger.Info(fmt.Sprintf("Writing to Topic - %v", topic))

	err := writer.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.FormatInt(time.Now().Unix(), 10)),
		// create an arbitrary message payload for the value
		Value: []byte(msg),
	})

	if err != nil {
		logger.Error("Writing to topic failed : ", err)
		return
	}

	logger.Info(fmt.Sprintf("Message written to topic - %v", topic))
}

package client

import (
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	instanceProducerKafka ProducerKafka
	onceProducerKafka    sync.Once
)

type ProducerKafka interface {
	Initialize() error
	SendMessage(payload []byte, topic string) error
}

type producerKafka struct{
	kafka *kafka.Producer
}

func GetInstanceKafka() ProducerKafka {
	onceProducerKafka.Do(func() {
		instanceProducerKafka = &producerKafka{}
	})
	return instanceProducerKafka
}

func (pk *producerKafka )Initialize() error{

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}
	pk.kafka = kafkaProducer

	// Delivery report handler for produced messages
	go func() {
		for e := range pk.kafka.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return nil

}

func (pk *producerKafka)SendMessage(payload []byte, topic string) error {
// payload is an json struct using marshal, so when I send to here I need to put on payload the model
// struct to an json, using json.Marshal()

	err := pk.kafka.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)

	if err != nil {
		return err
	}
	
	pk.kafka.Flush(15 * 1000)

	return nil

}

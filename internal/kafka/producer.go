package kafka

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/config"
)

func PushMessageToQueue(topic string, message []byte) error {
	brokers := []string{config.Envs.KafkaBrokers}

	//Create connection
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}
	defer producer.Close()

	//Create a new message
	m := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	//Send message
	partition, offset, err := producer.SendMessage(m)
	if err != nil {
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)

	return nil
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll 
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}
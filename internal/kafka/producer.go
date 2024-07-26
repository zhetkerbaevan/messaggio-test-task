package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/IBM/sarama"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/config"
)

func PushMessageToQueue(topic string, message []byte) error {
    brokers := []string{config.Envs.KafkaURL}

    // Create connection
    producer, err := ConnectProducer(brokers)
    if err != nil {
        return err
    }
    defer producer.Close()

    // Create a new message
    m := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }

    // Send message
    partition, offset, err := producer.SendMessage(m)
    if err != nil {
        return err
    }

    log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)

    return nil
}

func createTLSConfiguration(certFile, keyFile, caFile string) (*tls.Config, error) {
    // Load client cert
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return nil, err
    }

    // Load CA cert
    caCert, err := ioutil.ReadFile(caFile)
    if err != nil {
        return nil, err
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // Create TLS configuration
    return &tls.Config{
        Certificates:       []tls.Certificate{cert},
        RootCAs:            caCertPool,
        InsecureSkipVerify: true, // используйте false в продакшене
    }, nil
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
    tlsConfig, err := createTLSConfiguration("internal/config/certificates/client-cert.pem", "internal/config/certificates/client-key.pem", "internal/config/certificates/ca-cert.pem")
    if err != nil {
        return nil, err
    }

    config := sarama.NewConfig()
    config.Net.TLS.Enable = true
    config.Net.TLS.Config = tlsConfig
    config.Producer.Return.Successes = true
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5

    return sarama.NewSyncProducer(brokers, config)
}
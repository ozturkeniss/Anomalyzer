package eventstreamer

import (
	"log"

	"github.com/IBM/sarama"
)

type Publisher struct {
	producer sarama.SyncProducer
	topic    string
}

func NewPublisher(brokers []string, topic string) (*Publisher, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &Publisher{producer: producer, topic: topic}, nil
}

func (p *Publisher) Publish(message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(message),
	}
	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}
	return err
}

func (p *Publisher) Close() error {
	return p.producer.Close()
}

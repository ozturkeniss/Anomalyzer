package eventstreamer

import (
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer          sarama.Consumer
	topic             string
	partitionConsumer sarama.PartitionConsumer
}

func NewConsumer(brokers []string, topic string, groupID string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		consumer.Close()
		return nil, err
	}

	return &Consumer{
		consumer:          consumer,
		topic:             topic,
		partitionConsumer: partitionConsumer,
	}, nil
}

func (c *Consumer) Consume() ([]byte, error) {
	select {
	case msg := <-c.partitionConsumer.Messages():
		return msg.Value, nil
	case err := <-c.partitionConsumer.Errors():
		log.Printf("Error consuming message: %v", err)
		return nil, err
	}
}

func (c *Consumer) Close() error {
	if c.partitionConsumer != nil {
		c.partitionConsumer.Close()
	}
	if c.consumer != nil {
		return c.consumer.Close()
	}
	return nil
}

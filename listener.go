package bus

import (
	"errors"
	"github.com/Shopify/sarama"
)

var (
	// ErrTopicRequired is returned when topic is not passed as parameter in bus.ListenerConfig.
	ErrTopicRequired = errors.New("creating a new listener requires a non-empty topic")
	// ErrTopicRequired is returned when topic is not passed as parameter in bus.ListenerConfig.
	ErrHandlerFuncRequired = errors.New("creating a new listener requires a non-empty handler function")
)

type handlerFunc func(message *Message) error

// ListenerConfig carries the different variables to tune a newly started consumer,
// it exposes the same configuration available from official nsq go client.
type ListenerConfig struct {
	Topic       string
	Partition   int32
	Address     []string
	HandlerFunc handlerFunc
}

// On listen to a message from a specific topic using sarama consumer, returns
// an error if topic not passed or if an error ocurred while creating
// sarama consumer.
func On(lc ListenerConfig) error {
	if len(lc.Topic) == 0 {
		return ErrTopicRequired
	}

	if lc.HandlerFunc == nil {
		return ErrHandlerFuncRequired
	}

	address := lc.Address
	if len(address) == 0 {
		address = append(address, "localhost:9092")
	}

	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		return err
	}

	pc, err := consumer.ConsumePartition(lc.Topic, lc.Partition, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	go func(pc sarama.PartitionConsumer, handler handlerFunc) {
		for {
			select {
			case message := <-pc.Messages():
				m := Message{ConsumerMessage: message}
				handler(&m)
			}
		}
	}(pc, lc.HandlerFunc)

	return nil
}

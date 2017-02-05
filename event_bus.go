package eventbus

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"os"
)

var KAFKA_URL = os.Getenv("KAFKA_URL")

type EventBus interface {
	Emit(topic string, payload interface{}) error
	On(topic string, handler fnHandler) error
}

type Bus struct {
	Emitter  sarama.AsyncProducer
	Listener sarama.Consumer
}

type fnHandler func(payload []byte) error

func init() {
	if KAFKA_URL == "" {
		KAFKA_URL = "localhost:9093"
	}
}

func NewEventBus() (EventBus, error) {
	brokers := []string{KAFKA_URL}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Bus{producer, consumer}, nil
}

func (bus *Bus) Emit(topic string, payload interface{}) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}

	for {
		select {
		case bus.Emitter.Input() <- message:
		case <-bus.Emitter.Successes():
			return nil
		case err := <-bus.Emitter.Errors():
			return err
		}
	}
}

func (bus *Bus) On(topic string, handler fnHandler) error {
	partitions, err := bus.Listener.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		pc, err := bus.Listener.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			return err
		}

		go func() {
			for {
				select {
				case message := <-pc.Messages():
					handler(message.Value)
				}
			}
		}()
	}

	return nil
}

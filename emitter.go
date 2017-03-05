package bus

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

// Emitter exposes a interface for emitting and listening for events.
type Emitter interface {
	Emit(topic string, payload interface{}) error
	EmitAsync(topic string, payload interface{}) error
}

// EmitterConfig carries the different variables to tune a newly started emitter,
// it exposes the same configuration available from official nsq go client.
type EmitterConfig struct {
	Address       []string
	MaxRetry      int
	ReturnSuccess bool
	ReturnErrors  bool
}

type eventEmitter struct {
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
	address       []string
}

// NewEmitter returns a new eventEmitter configured with the
// variables from the config parameter, or returning an non-nil err
// if an error ocurred while creating kafka producer and async producer.
func NewEmitter(ec EmitterConfig) (emitter Emitter, err error) {
	config := newEmitterConfig(ec)

	address := ec.Address
	if len(address) == 0 {
		address = append(address, "localhost:9092")
	}

	syncProducer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		return
	}

	asyncProducer, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		return
	}

	emitter = &eventEmitter{syncProducer, asyncProducer, address}

	return
}

// Emit emits a message to a specific topic using sarama sync producer, returning
// an error if encoding payload fails or if an error ocurred while publishing
// the message.
func (ee eventEmitter) Emit(topic string, payload interface{}) (err error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(body),
	}

	_, _, err = ee.syncProducer.SendMessage(message)

	return
}

// Emit emits a message to a specific topic using sarama async producer, but does not wait for
// the response from `kafka`. Returns an error if encoding payload fails and
// logs to console if an error ocurred while publishing the message.
func (ee eventEmitter) EmitAsync(topic string, payload interface{}) (err error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(body),
	}

	ee.asyncProducer.Input() <- message

	go func(topic string) {
		for {
			select {
			case err := <-ee.asyncProducer.Errors():
				log.Println(err)
			}
		}
	}(topic)

	return
}

func newEmitterConfig(ec EmitterConfig) (config *sarama.Config) {
	config = sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal

	if ec.MaxRetry != 0 {
		config.Producer.Retry.Max = ec.MaxRetry
	}

	if ec.ReturnSuccess {
		config.Producer.Return.Successes = ec.ReturnSuccess
	}

	if ec.ReturnErrors {
		config.Producer.Return.Errors = ec.ReturnErrors
	}

	return
}

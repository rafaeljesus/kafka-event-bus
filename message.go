package bus

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

// Message carries sarama.ConsumerMessage fields and methods and
// adds extra fields for handling messages internally.
type Message struct {
	*sarama.ConsumerMessage
}

// DecodePayload deserializes data (as []byte) and creates a new struct passed by parameter.
func (m *Message) DecodePayload(v interface{}) (err error) {
	err = json.Unmarshal(m.Value, v)

	return
}

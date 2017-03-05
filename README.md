## Event Bus Kafka

* A tiny wrapper around [sarama](https://github.com/Shopify/sarama) topic and consumer.

## Installation
```bash
go get -u github.com/rafaeljesus/kafka-event-bus
```

## Usage
The kafka-event-bus package exposes a interface for emitting and listening events.

### Emitter
```go
import "github.com/rafaeljesus/kafka-event-bus"

topic := "events"
emitter, err := bus.NewEmitter(EmitterConfig{
  Address:  "localhost:9092",
  MaxRetry: 10,
})

e := event{}
if err = emitter.Emit(topic, &e); err != nil {
  // handle failure to emit message
}

// emitting messages on a async fashion
if err = emitter.EmitAsync(topic, &e); err != nil {
  // handle failure to emit message
}
```

### Listener
```go
import "github.com/rafaeljesus/kafka-event-bus"

if err = bus.On(bus.ListenerConfig{
  Topic:       "topic",
  HandlerFunc: handler,
  Partition:   0,
}); err != nil {
  // handle failure to listen a message
}

func handler(message *bus.Message) (err error) {
  e := event{}
  if err = message.DecodePayload(&e); err != nil {
    // handle failure to listen a message
    return
  }

  // handle message

  return
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Build Status](https://circleci.com/gh/rafaeljesus/kafka-event-bus.svg?style=svg)](https://circleci.com/gh/rafaeljesus/kafka-event-bus)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/kafka-event-bus)](https://goreportcard.com/report/github.com/rafaeljesus/kafka-event-bus)
[![Go Doc](https://godoc.org/github.com/rafaeljesus/kafka-event-bus?status.svg)](https://godoc.org/github.com/rafaeljesus/kafka-event-bus)

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Medium [@_jesus_rafael](https://medium.com/@_jesus_rafael) &nbsp;&middot;&nbsp;
> Twitter [@_jesus_rafael](https://twitter.com/_jesus_rafael)

## Event Bus Kafka

* A tiny wrapper around [sarama](https://github.com/Shopify/sarama) topic and consumer.

## Installation
```bash
go get -u https://github.com/rafaeljesus/kafka-event-bus
```

## Environment Variables
```bash
export KAFKA_URL=localhost:9093
```

## Usage
The kafka-event-bus package exposes a interface for emitting and listening events.

### Emitter
```go
import "github.com/rafaeljesus/kafka-event-bus"

topic := "events"
var event struct{}
eventBus, _ := eventbus.NewEventBus()

e := event{}
if err := eventBus.Emit(topic, &e); err != nil {
  // handle failure to emit message
}

```

### Listener
```go
import "github.com/rafaeljesus/kafka-event-bus"

topic := "events"
metricsChannel := "metrics"
notificationsChannel := "notifications"
eventBus, _ := eventbus.NewEventBus()

if err := eventBus.On(topic, metricsHandler); err != nil {
  // handle failure to listen a message
}

if err := eventBus.On(topic, notificationsHandler); err != nil {
  // handle failure to listen a message
}

func metricsHandler(payload []byte) (error) {
  e := event{}
  if err := json.Unmarshal(payload, &e); err != nil {
    // handle failure
  }
  // handle message
  return nil
}

func notificationsHandler(payload []byte) (interface{}, error) {
  e := event{}
  if err := json.Unmarshal(payload, &e); err != nil {
    // handle failure
  }
  // handle message
  return nil
}

```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/nsq-event-bus)](https://goreportcard.com/report/github.com/rafaeljesus/nsq-event-bus)

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Medium [@_jesus_rafael](https://medium.com/@_jesus_rafael) &nbsp;&middot;&nbsp;
> Twitter [@_jesus_rafael](https://twitter.com/_jesus_rafael)

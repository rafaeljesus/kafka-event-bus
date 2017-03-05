package bus

import (
	"sync"
	"testing"
)

func TestListenerOn(t *testing.T) {
	emitter, err := NewEmitter(EmitterConfig{})
	if err != nil {
		t.Errorf("Expected to initialize emitter %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	type event struct{ Name string }

	handler := func(message *Message) (err error) {
		e := event{}
		if err = message.DecodePayload(&e); err != nil {
			t.Errorf("Expected to unmarshal payload")
		}

		if e.Name != "event" {
			t.Errorf("Expected name to be equal event %s", e.Name)
		}

		wg.Done()
		return
	}

	if err = On(ListenerConfig{
		Topic:       "topic",
		Partition:   0,
		HandlerFunc: handler,
	}); err != nil {
		wg.Done()
		t.Errorf("Expected to listen a message %s", err)
	}

	e := event{"event"}
	if err := emitter.Emit("topic", &e); err != nil {
		wg.Done()
		t.Errorf("Expected to emit message %s", err)
	}

	wg.Wait()
}

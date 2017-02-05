package eventbus

import (
	"encoding/json"
	"testing"
	"time"
)

type event struct {
	Name string
}

func TestEventBusEmit(t *testing.T) {
	bus, err := NewEventBus()
	if err != nil {
		t.Errorf("Expected to initialize EventBus %s", err)
	}

	e := event{Name: "event"}
	if err := bus.Emit("topic", &e); err != nil {
		t.Errorf("Expected to emit message %s", err)
	}
}

func TestEventBusOn(t *testing.T) {
	bus, err := NewEventBus()
	if err != nil {
		t.Errorf("Expected to initialize EventBus %s", err)
	}

	handler := func(payload []byte) error {
		e := event{}
		if err := json.Unmarshal(payload, &e); err != nil {
			t.Errorf("Expected to unmarshal payload")
		}

		if e.Name != "event" {
			t.Errorf("Expected name to be equal event %s", e.Name)
		}

		return nil
	}

	e := event{Name: "event"}
	if err := bus.Emit("topic", &e); err != nil {
		t.Errorf("Expected to emit message %s", err)
	}

	if err := bus.On("topic", handler); err != nil {
		t.Errorf("Expected to listen a message %s", err)
	}

	time.Sleep(200 * time.Millisecond)
}

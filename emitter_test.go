package bus

import (
	"testing"
)

func TestNewEmitter(t *testing.T) {
	expected := "localhost:9092"

	emitter, err := NewEmitter(EmitterConfig{
		Address: []string{"localhost:9092"},
	})

	if err != nil {
		t.Errorf("Expected to initialize emitter %s", err)
	}

	e := emitter.(*eventEmitter)
	if e.address[0] != expected {
		t.Errorf("Expected emitter address %s, got %s", expected, e.address[0])
	}

	emitter, err = NewEmitter(EmitterConfig{})

	if err != nil {
		t.Errorf("Expected to initialize emitter %s", err)
	}

	e = emitter.(*eventEmitter)
	if e.address[0] != expected {
		t.Errorf("Expected emitter address %s, got %s", expected, e.address[0])
	}
}

func TestEmitterEmit(t *testing.T) {
	emitter, err := NewEmitter(EmitterConfig{})
	if err != nil {
		t.Errorf("Expected to initialize emitter %s", err)
	}

	type event struct{ Name string }
	e := event{"event"}

	if err := emitter.Emit("topic", &e); err != nil {
		t.Errorf("Expected to emit message %s", err)
	}
}

func TestEmitterEmitAsync(t *testing.T) {
	emitter, err := NewEmitter(EmitterConfig{})
	if err != nil {
		t.Errorf("Expected to initialize emitter %s", err)
	}

	type event struct{ Name string }
	e := event{"event"}

	if err := emitter.EmitAsync("topic", &e); err != nil {
		t.Errorf("Expected to emit message %s", err)
	}
}

package events

import (
	"testing"
	"time"
)

func TestAddEvent(t *testing.T) {
	name := "task 1"
	e := NewEvent(name)
	if e.name != name {
		t.Errorf("Output %q not equal to expected %q", e.name, name)
	}
	if e.frequency != time.Minute*0 {
		t.Errorf("Frequency should be zero value one-off tasks, got %s", e.frequency)
	}
	if e.recurring != false {
		t.Errorf("Recurring flag should be false, got %v", e.recurring)
	}
}

func TestAddEventRecurs(t *testing.T) {
	name := "task 1"
	frequency := time.Minute
	e := NewEventRecurs(name, frequency)
	if e.name != name {
		t.Errorf("Output %q not equal to expected %q", e.name, name)
	}
	if e.frequency != time.Minute {
		t.Errorf("Frequency does not match. Got %s, expected %s", e.frequency, frequency)
	}
	if e.recurring != true {
		t.Errorf("Recurring flag should be true, got %v", e.recurring)
	}
}

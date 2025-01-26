package events

import (
	"fmt"
	"time"
)

type EventDates struct {
	created   time.Time
	completed time.Time
	updated   time.Time
	reopen    time.Time
}

type Event struct {
	name      string
	dates     EventDates
	complete  bool
	recurring bool
	frequency time.Duration
}

func NewEvent(name string) *Event {
	e := Event{
		name: name,
		dates: EventDates{
			created: time.Now(),
		},
		complete:  false,
		recurring: false,
		frequency: time.Minute * 0,
	}
	return &e
}

func NewEventRecurs(name string, frequency time.Duration) Event {
	e := NewEvent(name)
	e.recurring = true
	e.frequency = frequency
	return *e
}

func (e *Event) Reopen() {
	if e.complete && e.recurring {
		if time.Since(e.dates.reopen) > 0 {
			e.complete = false
			e.dates.updated = time.Now()
		}
	}
}

func (e *Event) Complete() {
	if !e.complete {
		e.complete = true
		e.dates.updated = time.Now()
		e.dates.reopen = time.Now().Add(e.frequency)
	}
}

func (e Event) String() string {
	if e.recurring {
		return fmt.Sprintf("name: %s, created: %s, complete: %v, reopen: %s, updated: %s",
			e.name,
			e.dates.created.Format(time.UnixDate),
			e.complete,
			e.dates.reopen.Format(time.UnixDate),
			e.dates.updated.Format(time.UnixDate))
	}

	return fmt.Sprintf("name: %s, created: %s, complete: %v",
		e.name,
		e.dates.created.Format(time.UnixDate),
		e.complete)
}

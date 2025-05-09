package domain

import (
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"time"

	"github.com/google/uuid"
)

type Aggregate struct {
	id         string
	events     []*event.Event
	applyEvent func(*event.Event)
	version    int
}

func NewAggregate(id string, applyEventFunc func(ev *event.Event)) *Aggregate {
	if applyEventFunc == nil {
		applyEventFunc = func(ev *event.Event) {}
	}
	return &Aggregate{
		id:         id,
		events:     make([]*event.Event, 0),
		applyEvent: applyEventFunc,
	}
}

func (a *Aggregate) ID() string {
	return a.id
}

/*func (a *Aggregate) Name() string {
	return a.name
}*/

func (a *Aggregate) Version() int {
	return a.version
}

func (a *Aggregate) PendingVersion() int {
	return len(a.events) + a.version
}

func (a *Aggregate) Events() []*event.Event {
	return a.events
}

func (a *Aggregate) Commit() {
	a.version = a.PendingVersion()
	a.events = make([]*event.Event, 0)
}

func (a *Aggregate) ApplyEvent(event *event.Event) {
	a.applyEvent(event)
}

func (a *Aggregate) AddEvent(eventType event.EventType, data []byte) {
	ev := &event.Event{
		ID:          uuid.New(),
		Type:        eventType,
		Data:        data,
		Timestamp:   time.Now(),
		Version:     a.PendingVersion(),
		AggregateID: a.ID(),
	}
	a.events = append(a.events, ev)
	a.ApplyEvent(ev)
}

func (a *Aggregate) SetID(value string) {
	a.id = value
}

func (a *Aggregate) SetName(value string) {
	a.id = value
}

func (a *Aggregate) SetVersion(value int) {
	a.version = value
}

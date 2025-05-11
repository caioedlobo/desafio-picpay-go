package event

import (
	"github.com/google/uuid"
	"time"
)

const (
	UserCreated     EventType = "USER_CREATED"
	UserNameUpdated EventType = "USER_NAME_UPDATED"
	UserDeleted     EventType = "USER_DELETED"
)

type EventType string

type Event struct {
	ID          uuid.UUID `json:"id"`
	Type        EventType `json:"type"`
	Data        []byte    `json:"data"`
	Timestamp   time.Time `json:"timestamp"`
	Version     int       `json:"version"`
	AggregateID string    `json:"aggregate_id"`
}

type EventSourcedAggregate interface {
	ID() string
	Events() []*Event
	//Name() string
	Commit()
	ApplyEvent(*Event)

	SetID(string)
	SetName(string)
	SetVersion(int)
	NewEvent(eventType EventType, data []byte)
}

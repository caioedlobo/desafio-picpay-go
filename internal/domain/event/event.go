package event

import (
	"github.com/google/uuid"
	"time"
)

const (
	UserCreated EventType = "USER_CREATED"
	UserUpdated EventType = "USER_UPDATED"
	UserDeleted EventType = "USER_DELETED"
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

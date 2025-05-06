package event

import (
	"context"
)

type EventRepository interface {
	AppendEvent(ctx context.Context, ev []*Event) error
	GetEvents(ctx context.Context, aggregateID string) ([]*Event, error)
}

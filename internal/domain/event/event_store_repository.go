package event

import (
	"context"
)

type EventRepository interface {
	AppendEvent(ctx context.Context, ev []*Event) error
	Get(ctx context.Context, aggregateID string) (EventSourcedAggregate, error)
}

package event

import (
	"context"
)

type EventRepository interface {
	AppendEvent(ctx context.Context, ev []*Event) error
	Get(ctx context.Context, aggregateID string) (EventSourcedAggregate, error)
	Get2(ctx context.Context, aggregateID string, applyEvent func(ev *Event)) (EventSourcedAggregate, error)
}

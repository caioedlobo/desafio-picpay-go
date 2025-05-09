package eventstore

import (
	"context"
	"fmt"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEventStore struct {
	db *pgxpool.Pool
}

func NewPostgresEventStore(db *pgxpool.Pool) *PostgresEventStore {
	return &PostgresEventStore{
		db: db,
	}
}

func (s *PostgresEventStore) AppendEvent(ctx context.Context, ev []*event.Event) error {

	if len(ev) == 0 {
		return fmt.Errorf("no events to save")
	}
	eventsDuplicateMap := make(map[event.EventType]struct{})
	for _, v := range ev {
		if _, exists := eventsDuplicateMap[v.Type]; exists {
			return fmt.Errorf("duplicate event found: %v", v.Type)
		} else {
			eventsDuplicateMap[v.Type] = struct{}{}
		}
	}

	rows := make([][]any, len(ev))
	for i, ev := range ev {
		rows[i] = []any{
			ev.ID,
			ev.Type,
			ev.Data,
			ev.Timestamp,
			ev.Version,
			ev.AggregateID,
		}
	}

	_, err := s.db.CopyFrom(
		ctx,
		pgx.Identifier{"events"},
		[]string{"id", "type", "data", "timestamp", "version", "aggregate_id"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return fmt.Errorf("failed to copy events: %w", err)
	}
	return err
}

func (s *PostgresEventStore) Get(ctx context.Context, aggregateID string) (event.EventSourcedAggregate, error) {
	query := `
        SELECT id, type, data, timestamp, version, aggregate_id
        FROM events
        WHERE aggregate_id = $1
        ORDER BY version ASC
    `

	rows, err := s.db.Query(ctx, query, aggregateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	defer rows.Close()

	ag := domain.NewAggregate(aggregateID, nil)
	for rows.Next() {
		var ev event.Event

		err = rows.Scan(
			&ev.ID,
			&ev.Type,
			&ev.Data,
			&ev.Timestamp,
			&ev.Version,
			&ev.AggregateID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		ag.ApplyEvent(&ev)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}

	return ag, nil
}

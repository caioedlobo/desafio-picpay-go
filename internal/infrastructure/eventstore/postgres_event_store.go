package eventstore

import (
	"context"
	"database/sql"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
)

type PostgresEventStore struct {
	db *sql.DB
}

func NewPostgresEventStore(db *sql.DB) *PostgresEventStore {
	return &PostgresEventStore{
		db: db,
	}
}

func (s *PostgresEventStore) AppendEvent(ctx context.Context, event *event.Event) error {
	query := `
        INSERT INTO events (id, type, data, timestamp, version, aggregate_id)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	_, err := s.db.ExecContext(
		ctx,
		query,
		event.ID,
		event.Type,
		event.Data,
		event.Timestamp,
		event.Version,
		event.AggregateID,
	)

	return err
}

func (s *PostgresEventStore) GetEvents(ctx context.Context, aggregateID string) ([]*event.Event, error) {
	query := `
        SELECT id, type, data, timestamp, version, aggregate_id
        FROM events
        WHERE aggregate_id = $1
        ORDER BY version ASC
    `

	rows, err := s.db.QueryContext(ctx, query, aggregateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*event.Event

	for rows.Next() {
		var event event.Event

		if err := rows.Scan(
			&event.ID,
			&event.Type,
			&event.Data,
			&event.Timestamp,
			&event.Version,
			&event.AggregateID,
		); err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

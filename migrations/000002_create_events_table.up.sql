CREATE TABLE events (
    id UUID PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    data BYTEA NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    version INTEGER NOT NULL,
    aggregate_id VARCHAR(255) NOT NULL
);

CREATE INDEX idx_events_aggregate_id ON events(aggregate_id);
CREATE INDEX idx_events_type ON events(type);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
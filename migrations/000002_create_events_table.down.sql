DROP TABLE IF EXISTS events;
DROP INDEX IF EXISTS idx_events_aggregate_id;
DROP INDEX IF EXISTS idx_events_type;
DROP EXTENSION IF EXISTS "uuid-ossp";
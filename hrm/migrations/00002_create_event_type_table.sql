-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event_types (
	id BIGSERIAL,
	event_type_name  VARCHAR(50) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	UNIQUE(event_type_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event_types ;
-- +goose StatementEnd

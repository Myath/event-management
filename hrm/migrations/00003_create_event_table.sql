-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id BIGSERIAL,
    event_type_id INTEGER,
    event_name TEXT NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(50) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    published_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    
    PRIMARY KEY(id),
    CONSTRAINT fk_event_type
       FOREIGN KEY (event_type_id)
           REFERENCES event_types (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events ;
-- +goose StatementEnd

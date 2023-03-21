-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_events (
    user_id INTEGER,
    event_id INTEGER,
    status SMALLINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user_event_user
        FOREIGN KEY (user_id)
            REFERENCES users (id),
    CONSTRAINT fk_user_event_event
        FOREIGN KEY (event_id)
            REFERENCES events (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE links(
    id SERIAL PRIMARY KEY NOT NULL,
    short TEXT,
    source TEXT,
    user_id BIGINT REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOT()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd

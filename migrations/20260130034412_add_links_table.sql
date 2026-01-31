-- +goose Up
-- +goose StatementBegin
CREATE TABLE links(
    id SERIAL PRIMARY KEY NOT NULL,
    short TEXT UNIQUE NOT NULL,
    source TEXT NOT NULL, 
    user_id BIGINT REFERENCES users(id) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOT() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd

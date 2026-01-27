-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE,
    password TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_login_users ON users(login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP INDEX idx_login_users
-- +goose StatementEnd

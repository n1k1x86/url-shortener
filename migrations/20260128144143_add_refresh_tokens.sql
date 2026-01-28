-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_tokens(
    id SERIAL PRIMARY KEY,
    jti TEXT,
    hash TEXT,
    user_id BIGINT REFERENCES users(id),
    replaced_by BIGINT,
    is_revoked BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_jti ON refresh_tokens(jti);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
-- +goose StatementEnd

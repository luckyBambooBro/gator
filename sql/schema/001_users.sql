
-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at timestamp (2) NOT NULL,
    updated_at timestamp (2) NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
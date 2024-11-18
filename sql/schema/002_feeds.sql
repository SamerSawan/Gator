-- +goose Up
CREATE TABLE feeds (
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds CASCADE;
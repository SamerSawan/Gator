-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_url TEXT NOT NULL,
    CONSTRAINT fk_feed FOREIGN KEY (feed_url) REFERENCES feeds(url) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts CASCADE;

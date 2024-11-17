-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds
ON feeds.url = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT 
    *,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows 
INNER JOIN users
ON users.id = feed_follows.user_id
INNER JOIN feeds
ON feeds.url = feed_follows.feed_id
WHERE feed_follows.user_id = $1;
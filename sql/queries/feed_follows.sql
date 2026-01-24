-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
/*
The id is taken as filter and not feed_id because
the delete HTTP request usually doesn't include a body,
so we need to identify the feed_follow by its id.
*/
DELETE FROM feed_follows
WHERE id = $1 AND user_id = $2;

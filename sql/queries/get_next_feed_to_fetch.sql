-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;
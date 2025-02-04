-- name: PGAdvisoryLock :exec
SELECT pg_advisory_lock(@key);

-- name: PGAdvisoryUnlock :exec
SELECT pg_advisory_unlock(@key);
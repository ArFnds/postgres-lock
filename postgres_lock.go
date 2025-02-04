package postgreslock

import (
	"context"
	"hash/fnv"

	"github.com/ArFnds/postgres-lock/internal"
)

type PostgresDistributedLock struct {
	name    string
	queries *internal.Queries
	key     int64
}

func NewPostgresDistributedLock(name string, db internal.DBTX) *PostgresDistributedLock {
	return &PostgresDistributedLock{
		name:    name,
		queries: internal.New(db),
		key:     keyNameAsHash64(name),
	}
}

func (p *PostgresDistributedLock) Name() string {
	return p.name
}

func (p *PostgresDistributedLock) Acquire(ctx context.Context) error {
	return p.queries.PGAdvisoryLock(ctx, p.key)
}

func (p *PostgresDistributedLock) Release(ctx context.Context) error {
	return p.queries.PGAdvisoryUnlock(ctx, p.key)
}

var _ DistributedLock = (*PostgresDistributedLock)(nil)

// `pg_try_advisory_lock` takes a bigint rather than any kind of human-readable
// name. Just so we don't have to choose a random integer, hash a provided name
// to a bigint-compatible 64-bit uint64 and use that.
func keyNameAsHash64(keyName string) int64 {
	hash := fnv.New64()
	_, err := hash.Write([]byte(keyName))
	if err != nil {
		panic(err)
	}
	return int64(hash.Sum64())
}

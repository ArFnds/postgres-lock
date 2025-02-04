package postgreslock

import "context"

type DistributedLock interface {
	Name() string
	Acquire(ctx context.Context) error
	Release(ctx context.Context) error
}

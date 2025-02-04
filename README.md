# Postgres Distributed Lock

The postgres-lock package offers distributed synchronization lock based on [PostgreSQL advisory locks](https://www.postgresql.org/docs/9.4/explicit-locking.html#ADVISORY-LOCKS). For example:

```go
ctx := context.Background()
conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
if err != nil {
    t.Fatalf("Failed to connect %v", err)
}
defer conn.Close(ctx)

lock := postgreslock.NewPostgresDistributedLock("test", conn)

fmt.Println("waiting for lock")
err = lock.Acquire(ctx)
if err != nil {
    t.Fatalf("failed to acquire lock: %v", err)
}
defer func() {
    lock.Release(ctx)
    fmt.Println("lock released")
}()

fmt.Println("lock acquired")

time.Sleep(2 * time.Second)
```

## Implementation notes

Under the hood, [Postgres advisory locks can be based on either one 64-bit integer value or a pair of 32-bit integer values](https://www.postgresql.org/docs/12/functions-admin.html#FUNCTIONS-ADVISORY-LOCKS). Because of this, the lock take a name that will be hashed in a 64-bit int using FNV-1.

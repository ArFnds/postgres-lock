package postgreslock_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	_ "database/sql/driver"

	postgreslock "github.com/ArFnds/postgres-lock"
	"github.com/jackc/pgx/v5"
)

func TestPostgres(t *testing.T) {
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
}

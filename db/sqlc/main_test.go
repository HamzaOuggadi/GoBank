package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:hamza123@localhost:5432/go_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	testQueries = New(pool)

	os.Exit(m.Run())

}

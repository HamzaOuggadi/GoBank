package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:hamza123@localhost:5432/go_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}

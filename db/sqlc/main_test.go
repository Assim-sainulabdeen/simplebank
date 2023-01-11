package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB      *sql.DB
	testQueries *Queries
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:tech123@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}

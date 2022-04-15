package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const dbDriver = "postgres"

var dbSource = fmt.Sprintf("postgresql://%s:%s@localhost:%s/simple_bank?sslmode=disable", os.Getenv("SB_PGUSER"), os.Getenv("SB_PGPASS"), os.Getenv("SB_PGPORT"))

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

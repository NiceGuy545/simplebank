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
dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
    var err error

    // Connect to DB
    testDB, err = sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer testDB.Close()

    // Initialize queries
    testQueries = New(testDB)

    // Run tests
    code := m.Run()

    // Exit with proper status code
    os.Exit(code)
}

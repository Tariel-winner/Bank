package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
const (
	dbDriver="postgres"
	dbSource="postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	testDB, err := sql.Open(dbDriver,dbSource)
	if err!=nil{
		log.Fatal("cannot connect to db:",err)
	}

	testQueries=New(testDB)
	os.Exit(m.Run())

}
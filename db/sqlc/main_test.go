package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "myPassword"
	dbname   = "postgres"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	testDB, err = sql.Open(dbDriver, psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}

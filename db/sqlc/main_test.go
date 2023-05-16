package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stefan-vl/my-bank/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())

}

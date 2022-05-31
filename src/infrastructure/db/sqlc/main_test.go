package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/caiofernandes00/simple_bank/src/application/utils"
	db "github.com/caiofernandes00/simple_bank/src/infrastructure/db/sqlc"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB
var dbDriver, dbSource string

func GetDsn() (string, string) {
	dbDriver = "postgres"
	dbSource = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	return dbDriver, dbSource
}

func TestMain(m *testing.M) {
	var err error
	utils.LoadEnv()

	dbDriver, dbSource := GetDsn()
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)
	os.Exit(m.Run())
}

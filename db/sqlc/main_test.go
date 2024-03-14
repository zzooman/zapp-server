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
	dbSource = "postgresql://root:1033@localhost:5432/zapp?sslmode=disable"
)

var testStore *Store
var connPool *pgxpool.Pool

func TestMain(m *testing.M){
	var err error
	connPool, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)			
	}
		
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
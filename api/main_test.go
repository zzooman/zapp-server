package api

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/utils"
)



var testStore db.Store
var connPool *pgxpool.Pool

func TestMain(m *testing.M){	
	config, err := utils.LoadConfig("../")
	if err != nil { log.Fatal("cannot load config:", err)}
	connPool, err = pgxpool.New(context.Background(), config.DBSource)
	
	if err != nil { log.Fatal("cannot connect to db:", err) }
		
	testStore = db.NewStore(connPool)
	os.Exit(m.Run())
}
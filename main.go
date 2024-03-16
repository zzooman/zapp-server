package main

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zzooman/zapp-server/api"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "1033"
	dbname   = "zapp"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("%s://%s:%s@%s:%d/%s", dbDriver, user, password, host, port, dbname))
	if err != nil {
		panic(err)
	}		

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		panic(err)
	}
}	
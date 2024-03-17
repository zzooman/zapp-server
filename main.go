package main

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zzooman/zapp-server/api"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/utils"
)



func main() {
	config, err := utils.LoadConfig(".")
	if err != nil { panic(err) }

	conn, err := pgxpool.New(context.Background(), fmt.Sprintf(config.DBSource))
	if err != nil { panic(err) }		

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}	
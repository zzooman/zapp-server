package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zzooman/zapp-server/api"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/utils"
)



func main() {
	config, err := utils.LoadConfig(".")
	if err != nil { panic(err) }

	m, err := migrate.New("file://db/migration", config.DBSource)
	if err != nil { panic(err) }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Failed to run migrate up: %v\n", err)
    }
    fmt.Println("Migrations applied successfully!")

	conn, err := pgxpool.New(context.Background(), fmt.Sprintf(config.DBSource))
	if err != nil { panic(err) }		

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}	
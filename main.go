package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stefan-vl/my-bank/api"
	db "github.com/stefan-vl/my-bank/db/sqlc"
	"github.com/stefan-vl/my-bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot log config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("coul not connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}

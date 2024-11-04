package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/HuyDangDev/simplebank/api"
	db "github.com/HuyDangDev/simplebank/db/sqlc"
	"github.com/HuyDangDev/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannnot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start sever: ", err)
	}
}

package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"riddle_with_numbers/api"
	"riddle_with_numbers/db/sqlc"
	_ "riddle_with_numbers/docs"
	"riddle_with_numbers/util"
)

// @title: Riddle with numbers
func main() {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}
	dbCon, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(dbCon)

	server, err := api.NewServer(&conf, store)
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		return
	}
	err = server.Start(conf.ServerAddress)
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		return
	}
}

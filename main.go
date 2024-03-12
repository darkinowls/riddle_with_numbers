package main

import (
	"fmt"
	"riddle_with_numbers/api"
	_ "riddle_with_numbers/docs"
)

// @title: Riddle with numbers
func main() {
	server := api.NewServer()
	err := server.Router.Run(":8084")
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		return
	}
}

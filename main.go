package main

import (
	"riddle_with_numbers/api"
	_ "riddle_with_numbers/docs"
)

// @title: Riddle with numbers
func main() {
	server := api.NewServer()
	err := server.Router.Run(":8084")
	if err != nil {
		return
	}
}

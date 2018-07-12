package main

import (
	"github.com/defact/user/database"
	"github.com/defact/user/server"
)

func main() {
	database.Initialize()
	server.Start()
}

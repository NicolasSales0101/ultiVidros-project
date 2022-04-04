package main

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}

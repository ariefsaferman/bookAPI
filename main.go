package main

import (
	"bookAPI/db"
	"bookAPI/routers"
	"log"
)

func main() {
	var PORT = ":8080"

	// repo.Init()
	// routers.StartServer().Run(PORT)
	err := db.Connect()
	if err != nil {
		log.Println("failed to connect DB", err)
	}

	log.Println("Database Connected")

	routers.StartServer().Run(PORT)
}

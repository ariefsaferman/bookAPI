package main

import (
	repo "bookAPI/repository"
	"bookAPI/routers"
)

func main() {
	var PORT = ":8080"

	repo.Init()
	routers.StartServer().Run(PORT)
}

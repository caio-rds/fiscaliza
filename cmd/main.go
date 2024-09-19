package main

import (
	"fiscaliza/internal/api"
	"fiscaliza/internal/database"
)

func main() {

	var db = database.ConnectDB()

	router := api.NewRouter()
	router.StartRouter(db)
}

package main

import (
	"fiscaliza/internal/api"
	"fiscaliza/internal/services"
)

func main() {

	var db = services.ConnectDB()

	router := api.NewRouter()
	router.RouteOne(db)
}

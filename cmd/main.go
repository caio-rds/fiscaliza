package main

import (
	"community_voice/internal/api"
	"community_voice/internal/services"
)

func main() {

	var db = services.ConnectDB()

	router := api.NewRouter()
	router.RouteOne(db)
}

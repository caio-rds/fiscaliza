package main

import (
	"community_voice/internal/api"
	"community_voice/internal/models"
)

func main() {

	var db = models.ConnectDB()

	router := api.NewRouter()
	router.RouteOne(db)
}

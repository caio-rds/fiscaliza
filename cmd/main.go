package main

import (
	"community_voice/internal/api"
	"community_voice/internal/database"
)

func main() {

	var db = database.ConnectDB()

	router := api.NewRouter()
	router.RouteOne(db)
}

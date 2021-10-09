package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Subham-Panda/AppointyInstagramAPI/database"
	"github.com/Subham-Panda/AppointyInstagramAPI/routers"
)

func main() {
	port := ":8080"
	connection := database.ConnectDatabase()
	routers.HandleRoutes(connection)
	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
	defer database.DisconnectDatabase(connection)
}

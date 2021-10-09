package routers

import (
	"net/http"

	"github.com/Subham-Panda/AppointyInstagramAPI/controllers"
	"github.com/Subham-Panda/AppointyInstagramAPI/database"
)

func handleUsersRoutes(connection *database.DatabaseConnection) {
	usersApi := &controllers.UsersAPI{MongoDatabase: connection.Client.Database("appointyinstagram")}
	http.Handle("/users/", usersApi)
}
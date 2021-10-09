package routers

import "github.com/Subham-Panda/AppointyInstagramAPI/database"

func HandleRoutes(connection *database.DatabaseConnection) {
	handlePostsRoutes(connection)
	handleUsersRoutes(connection)
}
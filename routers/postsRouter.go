package routers

import (
	"net/http"

	"github.com/Subham-Panda/AppointyInstagramAPI/controllers"
	"github.com/Subham-Panda/AppointyInstagramAPI/database"
)

func handlePostsRoutes(connection *database.DatabaseConnection) {
	postsApi := &controllers.PostsAPI{MongoDatabase: connection.Client.Database("appointyinstagram")}
	http.Handle("/posts/", postsApi)
}
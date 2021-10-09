package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Subham-Panda/AppointyInstagramAPI/models"
	"github.com/Subham-Panda/AppointyInstagramAPI/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsAPI struct {
	MongoDatabase *mongo.Database
	ctx           context.Context
}

func (postsApi *PostsAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postsApi.get(w, r)
	case http.MethodPost:
		postsApi.post(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
	}
}

func (postsApi *PostsAPI) get(w http.ResponseWriter, r *http.Request) {
	postsCollection := postsApi.MongoDatabase.Collection("posts")
	if (len(strings.Split(r.URL.String(), "/")) == 4) && (strings.Split(r.URL.String(), "/")[1] == "posts") && (strings.Split(r.URL.String(), "/")[2] == "users") {
		id, err := utils.GetUserIDFromPostURL(r)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Valid User ID Required")
			return
		}
		primitiveObjectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Valid User ID Required")
			return
		}

		postsCursor, err := postsCollection.Find(postsApi.ctx, bson.M{"user": primitiveObjectID})
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "No user found with the given id")
		}
		var posts []bson.M
		err = postsCursor.All(postsApi.ctx, &posts)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Error finding posts of the given user")
		}
		if(len(posts) == 0) {
			utils.RespondWithError(w, http.StatusNotFound, "No posts of the user")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, posts)
	} else {
		id, err := utils.GetIDFromURL(r)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Valid Post ID Required")
			return
		}

		primitiveObjectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Valid Post ID Required")
			return
		}

		var post bson.M
		err = postsCollection.FindOne(postsApi.ctx, bson.M{"_id": primitiveObjectID}).Decode(&post)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Valid Post ID Required")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, post)
	}

}

func (postsApi *PostsAPI) post(w http.ResponseWriter, r *http.Request) {
	postsCollection := postsApi.MongoDatabase.Collection("posts")
	usersCollection := postsApi.MongoDatabase.Collection("users")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		utils.RespondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}

	if (!utils.CompareJSONToStruct(body, models.CheckPostRequestBody{})) {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user bson.M
	err = usersCollection.FindOne(postsApi.ctx, bson.M{"_id": post.User}).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}

	post.ID = primitive.NewObjectID()
	newPostResult, err := postsCollection.InsertOne(postsApi.ctx, post)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Error creating post")
	}
	fmt.Printf("Inserted 1 document into posts collection with id %v\n", newPostResult.InsertedID)
	utils.RespondWithJSON(w, http.StatusCreated, post)
}
package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Subham-Panda/AppointyInstagramAPI/models"
	"github.com/Subham-Panda/AppointyInstagramAPI/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersAPI struct {
	MongoDatabase *mongo.Database
	ctx           context.Context
}

func (usersApi *UsersAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		usersApi.get(w, r)
	case http.MethodPost:
		usersApi.post(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
	}
}
func (usersApi *UsersAPI) get(w http.ResponseWriter, r *http.Request) {
	usersCollection := usersApi.MongoDatabase.Collection("users")
	id, err := utils.GetIDFromURL(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Valid User ID Required")
		return
	}
	primitiveObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Valid User ID Required")
		return
	}
	var user bson.M
	err = usersCollection.FindOne(usersApi.ctx, bson.M{"_id": primitiveObjectID}).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Valid User ID Required")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}
func (usersApi *UsersAPI) post(w http.ResponseWriter, r *http.Request) {
	usersCollection := usersApi.MongoDatabase.Collection("users")

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
	if (!utils.CompareJSONToStruct(body, models.CheckUserRequestBody{})) {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.ID = primitive.NewObjectID()
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedPassword[:])
	newUserResult, err := usersCollection.InsertOne(usersApi.ctx, user)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Error creating user")
	}
	fmt.Printf("Inserted 1 document into users collection with id %v\n", newUserResult.InsertedID)
	utils.RespondWithJSON(w, http.StatusCreated, user)
}
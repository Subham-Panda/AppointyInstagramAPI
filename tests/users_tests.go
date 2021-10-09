package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/Subham-Panda/AppointyInstagramAPI/models"
)

func TestCreateUser(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"name":     "Subham Subhasish Panda",
		"password": "hello",
		"username": "iampanda",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://127.0.0.1:8080/users/", "application/json", responseBody)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	_, err = strconv.ParseBool(user.Username)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
}

func TestGetUser(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/users/6161d538485d398d971a2a18")
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	_, err = strconv.ParseBool(user.Username)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
}

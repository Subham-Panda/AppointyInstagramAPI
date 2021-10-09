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

func TestCreatePost(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"title":     "Title",
		"description": "description",
		"user": "6161d538485d398d971a2a18",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://127.0.0.1:8080/posts/", "application/json", responseBody)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	var post *models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	_, err = strconv.ParseBool(post.Title)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
}

func TestGetPost(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/posts/6161d538485d398d971a2l58")
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	var post *models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
	_, err = strconv.ParseBool(post.Title)
	if err != nil {
		t.Errorf("An Error Occured %v", err)
	}
}

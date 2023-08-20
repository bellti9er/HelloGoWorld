package controllers

import (
	"database/sql"
	"RESTful-API/models"
	"encoding/json"
	"net/http"
)

type PostController struct {
	DB *sql.DB
}

func NewPostController(db *sql.DB) *PostController {
	return &PostController{DB: db}
}

func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var p models.Post

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.DB.Exec("INSERT INTO posts (title, content) VALUES ($1, $2)", p.Title, p.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
package controllers

import (
	"database/sql"
	"encoding/json"
	"RESTful-API/models"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, name FROM users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	
	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.ID, &u.Name);
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _  := strconv.Atoi(params["userId"])
	row    := c.DB.QueryRow("SELECT id, name FROM users WHERE id=$1", id)

	var u models.User

	err := row.Scan(&u.ID, &u.Name);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}

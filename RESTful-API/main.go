package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"RESTful-API/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingController := controllers.NewPingController()
	userController := controllers.NewUserController(db)
	postController := controllers.NewPostController(db)

	router := mux.NewRouter()

	router.HandleFunc("/ping", pingController.Ping).Methods("GET")	
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", userController.GetUser).Methods("GET")
	router.HandleFunc("/posts", postController.CreatePost).Methods("POST")

	port := ":8080"
	log.Printf("🚀🚀🚀 Server Listening on port %s 🚀🚀🚀", port)
	log.Fatal(http.ListenAndServe(port, router))
}

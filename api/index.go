package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Person struct {
	gorm.Model

	Name string `json:"name"`
	Email string `json:"email"`
	Books []Book `json:"books"`
}

type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int `gorm:"unique_index"`
	PersonID int
}

type Server struct {
	*mux.Router
	db *gorm.DB
}

func NewServer() *Server {

	// load Env Variables
	HOST := os.Getenv("HOST")
	DB_PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("USER")
	NAME := os.Getenv("NAME")
	PASSWORD := os.Getenv("PASSWORD")

	// Data connection string
	DB_URI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", HOST, USER, NAME, PASSWORD, DB_PORT)
	
	// Open DB
	db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected successfully")
	}

	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})

	s := &Server{
		Router: mux.NewRouter(),
		db: db,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/people", func (w http.ResponseWriter, r *http.Request)  {
		var people []Person
		s.db.Find(&people)

		json.NewEncoder(w).Encode(people)
	}).Methods("GET")

	s.HandleFunc("/people/{ID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		res := s.db.Model(&Person{}).Where("id = ?", vars["ID"]).Update("email", vars["email"])

		if res.Error != nil {
			log.Fatal(res.Error)
		}

		w.Write([]byte("Success"))
	}).Queries("email", "{email}").Methods("PATCH")
}
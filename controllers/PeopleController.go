package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juddbaguio/rest-api-golang/interfaces"
	"github.com/juddbaguio/rest-api-golang/models"
	"gorm.io/gorm"
)

type PeopleController struct {
	interfaces.PeopleService
	DB *gorm.DB
}

func (controller *PeopleController) GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	controller.DB.Find(&people)
	json.NewEncoder(w).Encode(people)
}
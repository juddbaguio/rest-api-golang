package interfaces

import "github.com/juddbaguio/rest-api-golang/models"

type PeopleService interface {
	GetPeople() []models.Person
}


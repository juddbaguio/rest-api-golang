package routes

import (
	"github.com/gorilla/mux"
	"github.com/juddbaguio/rest-api-golang/controllers"
	"github.com/juddbaguio/rest-api-golang/service"
)

type PeopleHandler struct {
	controllers.PeopleController
}

func People(s *mux.Router) *mux.Router {
	controller := service.ServeService()
	peopleController := controller.InjectPeopleController()

	handler := &PeopleHandler{
		PeopleController: peopleController,
	}
	s.HandleFunc("/people", handler.GetPeople).Methods("GET")
	return s
}
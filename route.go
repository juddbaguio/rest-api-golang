package main

import (
	"sync"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/rest-api-golang/routes"
)


type IMuxRouter interface {
	NewServer() *mux.Router
}

type Server struct {
	*mux.Router
}

type router struct {}

var (
	m *router
	routerOnce sync.Once
)

func (router *router) NewServer() *mux.Router {
	s := mux.NewRouter()
	routes.People(s)
	return s
}

func MuxRouter() IMuxRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}

	return m
}
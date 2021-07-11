package service

import (
	"sync"

	"github.com/juddbaguio/rest-api-golang/controllers"
	"gorm.io/gorm"
)

type IServiceContainer interface {
	InjectPeopleController() controllers.PeopleController
}

type kernel struct {
	db *gorm.DB
}


func (k *kernel) InjectPeopleController() controllers.PeopleController {
	db := k.db

	return controllers.PeopleController{
		DB: db,
	}
}

var (
	k *kernel
	serveServiceOnce sync.Once
)

func ServeService() IServiceContainer {
	db := initializeDB()
	if k == nil {
		serveServiceOnce.Do(func() {
			k = &kernel{
				db: db,
			}
		})
	}

	return k
}
package context

import (
	"github.com/facebookgo/inject"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"practive_service/domain/service"
	"practive_service/domain/service/impl"
	"practive_service/infrastucture/controllers"
	"practive_service/infrastucture/repository"
	"practive_service/utils"
)

type Application struct {
	PersonController *controllers.PersonController `inject:""`
}

func BuildApplication(dbConn *gorm.DB) (*Application, error) {

	var app Application
	var g inject.Graph

	personController := controllers.PersonController{}
	personAppService := service.PersonAppService{}
	personServiceImpl := impl.PersonService{
		SelfAddress: utils.GetEnv("SELF_ADDRESS", "http://localhost:8080"),
	}

	personRepository := repository.CreatePersonRepository(dbConn)

	personServiceImpl.PersonRepository = personRepository
	personAppService.PersonService = personServiceImpl

	err := g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: &personController},
		&inject.Object{Value: &personAppService},
	)

	if err != nil {
		log.Fatalln(os.Stderr, err)
		os.Exit(1)
	}
	if err := g.Populate(); err != nil {
		log.Fatalln(os.Stderr, err)
		os.Exit(1)
	}

	return &app, err
}

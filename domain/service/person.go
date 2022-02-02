package service

import (
	"practive_service/domain/dto"
	"practive_service/domain/service/impl"
	"practive_service/infrastucture/models"
)

type PersonAppService struct {
	PersonService impl.PersonService
}

func (s PersonAppService) AddPerson(dto dto.Person) error {
	return s.AddPerson(dto)
}

func (s PersonAppService) FindPersonById(id int) (*models.Person, error) {
	return s.FindPersonById(id)
}

func (s PersonAppService) UpdatePerson(dto dto.Person) (*models.Person, error) {
	return s.UpdatePerson(dto)
}

func (s PersonAppService) DeletePerson(dto dto.Person) (*models.Person, error) {
	return s.DeletePerson(dto)
}

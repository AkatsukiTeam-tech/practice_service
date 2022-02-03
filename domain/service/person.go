package service

import (
	"practive_service/domain/dto"
	"practive_service/domain/service/impl"
	"practive_service/infrastucture/helpers"
	"practive_service/infrastucture/models"
)

type PersonAppService struct {
	PersonService impl.PersonService `inject:""`
}

func (s *PersonAppService) AddPerson(dto dto.Person) error {
	return s.PersonService.AddPerson(dto)
}

func (s *PersonAppService) FindPersonById(id int) (*models.Person, error) {
	return s.PersonService.FindPersonById(id)
}

func (s *PersonAppService) UpdatePerson(dto dto.Person) (*models.Person, error) {
	return s.PersonService.UpdatePerson(dto)
}

func (s *PersonAppService) DeletePerson(id int) error {
	return s.PersonService.DeletePerson(id)
}

func (s *PersonAppService) GetAllPerson(pagination helpers.Pagination) (*helpers.Pagination, error) {
	return s.PersonService.GetAllPerson(pagination)
}

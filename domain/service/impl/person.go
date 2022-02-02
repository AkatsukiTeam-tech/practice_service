package impl

import (
	"practive_service/domain/dto"
	"practive_service/infrastucture/models"
	"practive_service/infrastucture/repository"
	"time"
)

type PersonService struct {
	SelfAddress      string
	PersonRepository *repository.PersonRepository
}

func (s PersonService) AddPerson(dto dto.Person) error {
	person := &models.Person{
		ID:        dto.ID,
		FullName:  dto.FullName,
		Age:       dto.Age,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := s.PersonRepository.AddPerson(person)
	if err != nil {
		return err
	}

	return err
}

func (s PersonService) FindPersonById(id int) (*models.Person, error) {
	person, err := s.PersonRepository.FindPersonById(id)
	if err != nil {
		return nil, err
	}

	return person, err
}

func (s PersonService) UpdatePerson(dto dto.Person) (*models.Person, error) {
	person := &models.Person{
		ID:        dto.ID,
		FullName:  dto.FullName,
		Age:       dto.Age,
		CreatedAt: dto.CreatedAt,
		UpdateAt:  time.Time{},
	}
	person, err := s.PersonRepository.UpdatePerson(person)
	if err != nil {
		return nil, err
	}

	return person, err
}

func (s PersonService) DeletePerson(dto dto.Person) (*models.Person, error) {
	person := &models.Person{
		ID:        dto.ID,
		FullName:  dto.FullName,
		Age:       dto.Age,
		CreatedAt: dto.CreatedAt,
		UpdateAt:  dto.UpdateAt,
	}
	person, err := s.PersonRepository.DeletePerson(person)
	if err != nil {
		return nil, err
	}

	return person, err
}

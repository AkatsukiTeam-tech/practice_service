package impl

import (
	"log"
	"practive_service/domain/dto"
	"practive_service/infrastucture/helpers"
	"practive_service/infrastucture/models"
	"practive_service/infrastucture/repository"
	"time"
)

type PersonService struct {
	SelfAddress      string
	PersonRepository *repository.PersonRepository
}

func (s *PersonService) AddPerson(dto dto.Person) error {
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

func (s *PersonService) FindPersonById(id int) (*models.Person, error) {
	person, err := s.PersonRepository.FindPersonById(id)
	if err != nil {
		return nil, err
	}

	return person, err
}

func (s *PersonService) UpdatePerson(dto dto.Person) (*models.Person, error) {
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

func (s *PersonService) DeletePerson(id int) error {
	err := s.PersonRepository.DeletePerson(id)
	if err != nil {
		return err
	}

	return err
}

func (s *PersonService) GetAllPerson(pagination helpers.Pagination) (*helpers.Pagination, error) {
	persons, err := s.PersonRepository.GetAllPerson(&pagination)
	if err != nil {
		log.Printf("error while get persons: %s", err.Error())
		return &pagination, err
	}

	var dtoList = make([]dto.Person, 0)
	for _, person := range persons {
		personDto := dto.Person{
			ID:        person.ID,
			FullName:  person.FullName,
			Age:       person.Age,
			CreatedAt: person.CreatedAt,
			UpdateAt:  person.UpdateAt,
		}

		dtoList = append(dtoList, personDto)
	}
	pagination.Rows = dtoList

	return &pagination, err
}

func (s *PersonService) GetAllPersonByQuery(name string, age int) (*[]dto.Person, error) {
	persons, err := s.PersonRepository.GetAllPersonByQuery(name, age)
	if err != nil {
		log.Printf("error while get persons: %s", err.Error())
		return nil, err
	}

	var dtoList = make([]dto.Person, 0)
	for _, person := range persons {
		personDto := dto.Person{
			ID:        person.ID,
			FullName:  person.FullName,
			Age:       person.Age,
			CreatedAt: person.CreatedAt,
			UpdateAt:  person.UpdateAt,
		}

		dtoList = append(dtoList, personDto)
	}

	return &dtoList, err
}

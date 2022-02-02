package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"practive_service/infrastucture/models"
	"time"
)

type PersonRepository struct {
	db *gorm.DB
}

func CreatePersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) AddPerson(n *models.Person) error {
	return r.db.Create(&n).Error
}

func (r PersonRepository) FindPersonById(id int) (*models.Person, error) {
	res := &models.Person{}
	err := r.db.First(res, "id = ?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Person Not Found")
	}
	return res, nil
}

func (r PersonRepository) UpdatePerson(n *models.Person) (*models.Person, error) {
	res := &models.Person{}
	err := r.db.First(res, "id = ?", res.ID).UpdateColumns(
		map[string]interface{}{
			"full_name": n.FullName,
			"age":       n.Age,
			"update_at": time.Now(),
		}).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r PersonRepository) DeletePerson(n *models.Person) (*models.Person, error) {
	res := &models.Person{}
	err := r.db.First(res, "id = ?", res.ID).Delete(res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

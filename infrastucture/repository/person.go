package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"practive_service/infrastucture/helpers"
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
	err := r.db.First(res, "id = ?", n.ID).UpdateColumns(
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

func (r PersonRepository) DeletePerson(id int) error {
	err := r.db.Delete(models.Person{}, id).Error
	if err != nil {
		return err
	}
	return err
}

func (r PersonRepository) GetAllPerson(pagination *helpers.Pagination) ([]*models.Person, error) {
	var persons []*models.Person
	db := r.db.Model(persons)
	err := r.db.Scopes(helpers.Paginate(pagination, db)).Find(&persons).Error

	return persons, err
}

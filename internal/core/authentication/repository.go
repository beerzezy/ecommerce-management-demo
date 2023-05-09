package authentication

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Registration(request RequestRegistration) (interface{}, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) Registration(request RequestRegistration) (interface{}, error) {
	var result RequestRegistration
	if err := r.db.Table("account").Where("id = ?").Order("id").First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

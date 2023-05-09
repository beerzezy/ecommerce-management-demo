package usermanagement

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type Repository interface {
	GetAccountInfo(email string) (ResponseGetAccountInfo, error)
	CreateAccount(entity Accounts) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAccountInfo(email string) (ResponseGetAccountInfo, error) {
	var result ResponseGetAccountInfo

	sqlStr := `
	select account_id
	,concat(first_name, ' ', last_name)as full_name
	,role
	,email
	,hash_text
	from accounts
	where email = ?
	`
	if err := r.db.Raw(sqlStr, email).First(&result).Error; err != nil {
		log.Error("GetAccountInfo.repo Error")
		return result, err
	}

	return result, nil
}

func (r repository) CreateAccount(entity Accounts) error {
	tx := r.db.Begin()
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

package usermanagement

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type Service interface {
	GetAccountInfo(email string) (ResponseGetAccountInfo, error)
	CreateAccount(request RequestCreateAccount) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) GetAccountInfo(email string) (ResponseGetAccountInfo, error) {
	return s.repository.GetAccountInfo(email)
}

func (s service) CreateAccount(request RequestCreateAccount) error {
	//check account
	_, err := s.repository.GetAccountInfo(request.Email)
	if !gorm.IsRecordNotFoundError(err) {
		log.Error("This account already exists")
		return fmt.Errorf("CreateAccount Error: account already exists")
	}

	entity := Accounts{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Role:      1,
		HashText:  hashPassword([]byte(request.Password)),
		OpenDate:  time.Now(),
	}

	return s.repository.CreateAccount(entity)
}

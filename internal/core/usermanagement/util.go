package usermanagement

import (
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}

	return string(hash)
}

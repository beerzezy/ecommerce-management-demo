package authentication

import (
	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role int    `json:"role"`
	jwt.RegisteredClaims
}

type RequestRegistration struct {
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RespRegisAccount struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    ResponseRegistration `json:"data"`
}

type ResponseRegistration struct {
	AccountId int    `json:"account_id"`
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestGetAccountInfo struct {
	Email string `json:"email"`
}

type ResponseLogin struct {
	AccountId int    `json:"account_id"`
	FullName  string `json:"full_name"`
	Role      int    `json:"role"`
	Token     string `json:"token"`
}

type RespAccount struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

type Account struct {
	AccountId int    `json:"account_id"`
	FullName  string `json:"full_name"`
	Role      int    `json:"role"`
	Email     string `json:"email"`
	HashText  string `json:"hash_text"`
}

package usermanagement

import "time"

type RequestGetAccountInfo struct {
	Email string `json:"Email"`
}

type ResponseGetAccountInfo struct {
	AccountId int    `json:"account_id"`
	FullName  string `json:"full_name"`
	Role      int    `json:"role"`
	Email     string `json:"email"`
	HashText  string `json:"hash_text"`
}

type RequestCreateAccount struct {
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"Password"`
}

type Accounts struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      int       `json:"role"`
	Email     string    `json:"email"`
	HashText  string    `json:"hash_text"`
	OpenDate  time.Time `json:"open_date"`
}

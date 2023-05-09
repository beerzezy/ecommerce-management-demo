package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Service interface {
	Login(request RequestLogin) (ResponseLogin, error)
	Registration(request RequestRegistration) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) Login(request RequestLogin) (ResponseLogin, error) {

	resp, err := s.apiGetAccount(RequestGetAccountInfo{
		Email: request.Email,
	})
	if err != nil {
		return ResponseLogin{}, err
	}

	var account RespAccount
	if err := json.Unmarshal(resp, &account); err != nil {
		return ResponseLogin{}, err
	}

	matchEmail := false
	if request.Email == account.Data.Email {
		matchEmail = true
	}

	matchPwd := comparePasswords(account.Data.HashText, []byte(request.Password))
	if !matchPwd || !matchEmail {
		return ResponseLogin{}, fmt.Errorf("user or password invalid")
	}

	maxHour, err := strconv.Atoi(viper.GetString("jwt.maxHour"))
	if err != nil {
		return ResponseLogin{}, err
	}

	claims := &jwtCustomClaims{
		Name: account.Data.FullName,
		Role: account.Data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(maxHour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := viper.GetString("jwt.signingkey")
	signedJWT, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ResponseLogin{}, err
	}

	return ResponseLogin{
		AccountId: account.Data.AccountId,
		FullName:  account.Data.FullName,
		Role:      account.Data.Role,
		Token:     signedJWT,
	}, nil
}

func (s service) Registration(request RequestRegistration) error {
	resp, err := s.apiCreateAccount(request)
	var account RespAccount
	if err = json.Unmarshal(resp, &account); err != nil {
		return err
	}

	if account.Code != 0 {
		return fmt.Errorf(account.Message)
	}

	return nil
}

func (s service) apiGetAccount(data interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(data)
	fmt.Println("reqBody ===> ", string(reqBody))
	if err != nil {
		//log.Error(err)
		return []byte{}, err
	}

	subPrefix := "/account-info"
	usermanagementApi := viper.GetString("service.usermanagement") + subPrefix
	req, err := http.NewRequest("POST", usermanagementApi, bytes.NewBuffer(reqBody))
	if err != nil {
		//log.Error(err)
		return []byte{}, err
	}

	//req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-request-from", "authentication-service")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("err====>", err)
		//log.Error(err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err====>", err)
		//log.Error(err)
		return []byte{}, err
	}
	fmt.Println("-------------->")
	fmt.Println("------ usermanagement service ------>")
	//log.Debug(string(body))
	fmt.Println("------ usermanagement service ------>")
	fmt.Println("-------------->")

	return body, nil
}

func (s service) apiCreateAccount(data interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(data)
	fmt.Println("reqBody ===> ", string(reqBody))
	if err != nil {
		//log.Error(err)
		return []byte{}, err
	}

	subPrefix := "/registration"
	usermanagementApi := viper.GetString("service.usermanagement") + subPrefix
	req, err := http.NewRequest("POST", usermanagementApi, bytes.NewBuffer(reqBody))
	if err != nil {
		//log.Error(err)
		return []byte{}, err
	}

	//req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-request-from", "authentication-service")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("err====>", err)
		//log.Error(err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err====>", err)
		//log.Error(err)
		return []byte{}, err
	}
	fmt.Println("-------------->")
	fmt.Println("------ User Management Service ------>")
	//log.Debug(string(body))
	fmt.Println("------ User Management Service ------>")
	fmt.Println("-------------->")

	return body, nil
}

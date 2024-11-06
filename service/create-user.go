package service

import (
	"day-22/model"
	"day-22/repository"
	"errors"
	"fmt"
)

type UserService struct {
	RepoUser repository.UserRepositoryDB
}

func NewUserService(repo repository.UserRepositoryDB) *UserService {
	return &UserService{RepoUser: repo}
}

func (us *UserService) InputDataUser(name string, username string, password string) (map[string]interface{}, error) {
	if name == "" {
		return nil, errors.New("username tidak boleh kosong")
	}
	if username == "" {
		return nil, errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return nil, errors.New("password tidak boleh kosong")
	}

	user := model.User{
		Name:     name,
		Username: username,
		Password: password,
		Status:   "inActive",
		Token:    "nil",
	}

	err := us.RepoUser.CreateDataUser(user)
	if err != nil {
		fmt.Println("Error :", err)
	}

	response := map[string]interface{}{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
		"status":   user.Status,
		"token":    user.Token,
	}

	fmt.Println("berhasil input data user dengan id ", user.ID)

	return response, nil
}

func (us *UserService) CheckToken(token string) string {
	result := us.RepoUser.CheckToken(token)
	return result
}

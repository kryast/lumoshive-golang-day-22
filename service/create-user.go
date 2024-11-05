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

func (us *UserService) InputDataUser(name string, username string, password string) error {
	if name == "" {
		return errors.New("username tidak boleh kosong")
	}
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	user := model.User{
		Name:     name,
		Username: username,
		Password: password,
	}

	err := us.RepoUser.CreateDataUser(user)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data user dengan id ", user.ID)

	return nil
}

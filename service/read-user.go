package service

import "day-22/model"

func (us *UserService) GetAllDataUser() (*[]model.User, error) {
	user, err := us.RepoUser.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil

}

package service

import (
	"day-22/model"
	"fmt"
)

func (us *UserService) GetUserDetails(id int) (*model.User, error) {
	user, err := us.RepoUser.GetDetailUsers(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user details: %v", err)
	}
	return user, nil
}

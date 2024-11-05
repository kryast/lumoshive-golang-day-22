package service

import (
	"fmt"
)

func (us *UserService) GetAllDataUser() ([]map[string]interface{}, error) {

	users, err := us.RepoUser.GetAll()

	if err != nil {
		return nil, err
	}

	if users == nil || len(*users) == 0 {
		fmt.Println("No user found")
		return nil, err
	}

	var listUsers []map[string]interface{}
	for _, user := range *users {
		listUsers = append(listUsers, map[string]interface{}{
			"name":   user.Name,
			"status": user.Status,
		})
	}

	return listUsers, nil

}

package service

import (
	"day-22/model"
	"fmt"

	"github.com/google/uuid" // Import UUID untuk generate token
)

// LoginUser memverifikasi username dan password, lalu menghasilkan token baru dan mengubah status menjadi Active
func (us *UserService) LoginUser(user model.User) (model.User, error) {

	// Retrieve the user from the repository
	users, err := us.RepoUser.GetUserLogin(user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to retrieve user for login: %w", err)
	}

	// Generate a new token using UUID
	token := uuid.New().String()

	// Update the user's token and status
	users.Token = token
	users.Status = "Active" // Set status to "Active"

	// Save the updated user record in the repository
	err = us.RepoUser.UpdateUser(users)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to update user token and status: %w", err)
	}

	// Return the login response with user details and the new token
	response := model.User{
		ID:       users.ID,
		Name:     users.Name,
		Username: users.Username,
		Status:   users.Status,
		Token:    users.Token, // New token
	}

	return response, nil
}

package service

import (
	"day-22/model"
	"fmt"

	"github.com/google/uuid" // Import UUID untuk generate token
)

// LoginUser memverifikasi username dan password, lalu menghasilkan token baru dan mengubah status menjadi Active
func (us *UserService) LoginUser(user model.User) (map[string]interface{}, error) {

	users, err := us.RepoUser.GetUserLogin(user)

	if err != nil {
		return nil, err
	}
	// 3. Generate token baru menggunakan UUID
	token := uuid.New().String()

	// 4. Update token dan status pada user yang login
	users.Token = token
	users.Status = "Active" // Set status menjadi "Active"

	// 5. Update user record dengan token baru dan status Active
	err = us.RepoUser.UpdateUser(users)
	if err != nil {
		return nil, fmt.Errorf("failed to update user token and status: %v", err)
	}

	// 6. Kirim response dengan detail user dan token yang baru
	response := map[string]interface{}{
		"id":       users.ID,
		"name":     users.Name,
		"username": users.Username,
		"status":   users.Status,
		"token":    users.Token, // Token baru yang di-generate
	}

	return response, nil
}

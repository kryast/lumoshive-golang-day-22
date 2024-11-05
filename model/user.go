package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Token    string `json:"token"`
}

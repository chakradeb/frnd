package models

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Password string `json:"password"`
}

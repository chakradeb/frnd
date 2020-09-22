package models

type Profile struct {
	Username string `json:"username"`
	Followers int `json:"followers"`
}

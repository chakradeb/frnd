package models

type Session struct {
	Username string `json:"username"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

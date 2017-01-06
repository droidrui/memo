package model

type LoginInfo struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         interface{} `json:"user,omitempty"`
}

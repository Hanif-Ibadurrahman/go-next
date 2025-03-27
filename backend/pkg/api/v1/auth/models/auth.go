package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ResponseAuthLogin struct {
	AccessToken  string `json:"access_token"`
	ExpiredAt    int    `json:"expired_at"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type RequestAuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestCredentialValidate struct {
	Username string
	Level    int
	UserId   int
	Name     string
}

type ResponseAuthRefreshToken struct {
	AccessToken  string `json:"access_token"`
	ExpiredAt    int    `json:"expired_at"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func (v *RequestAuthLogin) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Username, validation.Required),
		validation.Field(&v.Password, validation.Required),
	)
}

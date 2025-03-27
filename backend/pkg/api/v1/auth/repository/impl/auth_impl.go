package impl

import (
	"backend/pkg/api/v1/auth/models"
	userModel "backend/pkg/api/v1/user/models"
	"context"

	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) Login(ctx context.Context, req models.RequestAuthLogin) (userModel.User, error) {
	var user userModel.User
	err := a.db.WithContext(ctx).
		Raw("SELECT id_user, username, name, level FROM users WHERE username = ? and password = ? and is_active = 1", req.Username, req.Password).
		Scan(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil
}

func (a *Auth) GetCredential(req models.RequestCredentialValidate) (string, error) {
	var result models.RequestCredentialValidate
	err := a.db.
		Raw("SELECT id_user, username, name, level FROM users WHERE username = ? and is_active = 1", req.Username).
		Scan(&result).Error

	if err != nil {
		return "", err
	}

	return result.Username, nil
}

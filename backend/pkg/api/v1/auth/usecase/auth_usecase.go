package usecase

import (
	"backend/app/config/constant"
	"backend/app/middleware/jwt"
	"backend/pkg/api/v1/auth/models"
	"context"
	"crypto/sha256"
	"encoding/hex"
)

func (a *Auth) LoginAuth(ctx context.Context, request models.RequestAuthLogin) (models.ResponseAuthLogin, error) {
	var login models.RequestAuthLogin
	hashed_password := ""
	for i := 0; i < 10; i++ {
		hash := sha256.Sum256([]byte(request.Password))
		hashed_password = hex.EncodeToString(hash[:])
	}

	login = models.RequestAuthLogin{
		Username: request.Username,
		Password: hashed_password,
	}

	user, err := a.repository.auth.Login(ctx, login)
	if err != nil {
		return models.ResponseAuthLogin{}, constant.ErrInternalServerError
	}

	accessToken, expiresAt, err := jwt.GenerateAccessToken(user.Username, int64(user.IDUser), user.Name)
	if err != nil {
		return models.ResponseAuthLogin{}, constant.ErrInternalServerError
	}

	refreshToken, _, err := jwt.GenerateRefreshToken(user.Username, int64(user.IDUser), user.Name)
	if err != nil {
		return models.ResponseAuthLogin{}, constant.ErrInternalServerError
	}

	return models.ResponseAuthLogin{
		AccessToken:  accessToken,
		ExpiredAt:    int(expiresAt),
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}, nil
}

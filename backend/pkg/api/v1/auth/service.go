package auth

import (
	"backend/pkg/api/v1/auth/models"
	"context"
)

type Service interface {
	LoginAuth(ctx context.Context, req models.RequestAuthLogin) (models.ResponseAuthLogin, error)
	// RefreshTokenAuth(username string) (models.ResponseAuthRefreshToken, error)
}

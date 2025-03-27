package repository

import (
	"backend/pkg/api/v1/auth/models"
	userModel "backend/pkg/api/v1/user/models"
	"context"
)

type AuthRepo interface {
	Login(ctx context.Context, request models.RequestAuthLogin) (userModel.User, error)
	GetCredential(models.RequestCredentialValidate) (string, error)
}

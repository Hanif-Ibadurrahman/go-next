package user

import (
	"backend/pkg/api/v1/user/models"
	"context"
)

type Service interface {
	Search(ctx context.Context, req models.QuerySearch) ([]models.UserDetail, error)
	CreateUser(ctx context.Context, req models.RequestCreateUser) (models.ResponseCreateUser, error)
	UpdateUser(ctx context.Context, req models.UpdateUser) error
	DeleteUser(ctx context.Context, userId int) error
}

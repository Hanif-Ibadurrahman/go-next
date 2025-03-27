package repository

import (
	"backend/pkg/api/v1/user/models"
	"context"
)

type UserRepo interface {
	Search(ctx context.Context, req models.QuerySearch) ([]models.UserDetail, error)
	UsernameExists(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, req models.User) (*models.User, error)
	UpdateUser(ctx context.Context, req models.User) error
	DeleteUser(ctx context.Context, userId int) error
}

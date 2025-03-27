package usecase

import (
	"backend/app/config/constant"
	"backend/pkg/api/v1/user/models"
	"context"
	"crypto/sha256"
	"encoding/hex"
)

func (u *User) Search(ctx context.Context, request models.QuerySearch) ([]models.UserDetail, error) {
	users, err := u.repository.user.Search(ctx, request)
	if err != nil {
		return nil, constant.ErrInternalServerError
	}

	return users, nil
}

func (u *User) CreateUser(ctx context.Context, request models.RequestCreateUser) (models.ResponseCreateUser, error) {
	// Add hashed password
	hashed_password := ""
	for i := 0; i < 10; i++ {
		hash := sha256.Sum256([]byte(request.Password))
		hashed_password = hex.EncodeToString(hash[:])
	}

	newUser := models.User{
		Level:    request.Level,
		Name:     request.Name,
		Username: request.Username,
		Password: hashed_password,
		IsActive: 1,
	}

	exists, err := u.repository.user.UsernameExists(ctx, request.Username)
	if err != nil {
		return models.ResponseCreateUser{}, constant.ErrInternalServerError
	}

	if exists {
		return models.ResponseCreateUser{}, constant.ErrUsernameAlreadyExists
	}

	user, err := u.repository.user.CreateUser(ctx, newUser)
	result := models.ResponseCreateUser{
		IDUser:   user.IDUser,
		Username: user.Username,
		Name:     user.Name,
		Level:    user.Level,
	}
	if err != nil {
		return models.ResponseCreateUser{}, constant.ErrInternalServerError
	}

	return result, nil
}

func (u *User) UpdateUser(ctx context.Context, request models.UpdateUser) error {
	var updateUser models.User
	hashed_password := ""
	if request.Password != nil {
		for i := 0; i < 10; i++ {
			hash := sha256.Sum256([]byte(*request.Password))
			hashed_password = hex.EncodeToString(hash[:])
		}

		updateUser = models.User{
			IDUser:   request.IDUser,
			Password: hashed_password,
		}
	} else if request.Password != nil && request.Name != nil {
		updateUser = models.User{
			IDUser:   request.IDUser,
			Password: hashed_password,
			Name:     *request.Name,
		}
	}

	if request.Name != nil {
		updateUser = models.User{
			IDUser: request.IDUser,
			Name:   *request.Name,
		}
	}

	err := u.repository.user.UpdateUser(ctx, updateUser)
	if err != nil {
		return constant.ErrInternalServerError
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, userId int) error {
	err := u.repository.user.DeleteUser(ctx, userId)
	if err != nil {
		return constant.ErrInternalServerError
	}

	return nil
}

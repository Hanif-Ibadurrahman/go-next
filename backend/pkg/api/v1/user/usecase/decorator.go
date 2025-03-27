package usecase

import (
	user "backend/pkg/api/v1/user/repository"
	userImpl "backend/pkg/api/v1/user/repository/impl"

	"gorm.io/gorm"
)

type Repository struct {
	user user.UserRepo
}

type User struct {
	repository Repository
}

func NewUserRepository(db *gorm.DB) Repository {
	return Repository{
		user: userImpl.NewUser(db),
	}
}

func New(repository Repository) *User {
	return &User{
		repository: repository,
	}
}

func Initialize(db *gorm.DB) *User {
	return New(NewUserRepository(db))
}

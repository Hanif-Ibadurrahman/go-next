package usecase

import (
	auth "backend/pkg/api/v1/auth/repository"
	authImpl "backend/pkg/api/v1/auth/repository/impl"

	"gorm.io/gorm"
)

type Repository struct {
	auth auth.AuthRepo
}

type Auth struct {
	repository Repository
}

func NewAuthRepository(db *gorm.DB) Repository {
	return Repository{
		auth: authImpl.NewAuth(db),
	}
}

func New(repository Repository) *Auth {
	return &Auth{
		repository: repository,
	}
}

func Initialize(db *gorm.DB) *Auth {
	return New(NewAuthRepository(db))
}

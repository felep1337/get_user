package use_case

import (
	"http_auth/internal/domain"
)

type usersRepository interface {
	FindByUUID(ID int) (*domain.User, error)
}

type UserRepo struct {
	usersRepository usersRepository
}

func NewUserRepo(data usersRepository) *UserRepo {
	return &UserRepo{usersRepository: data}
}

func (UR *UserRepo) FindByUUID(ID int) (*domain.User, error) {
	return UR.usersRepository.FindByUUID(ID)
}

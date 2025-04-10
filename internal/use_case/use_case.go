package use_case

import (
	"http_auth/internal/domain"
)

type UserRepo struct {
	Data domain.FindByUUID
}

func (UR *UserRepo) FindByUUID(ID int) (*domain.User, error) {
	return UR.Data.FindByUUID(ID)
}

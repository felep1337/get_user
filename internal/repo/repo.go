package repo

import (
	"errors"
	"http_auth/internal/domain"
)

type UserRepoMap struct {
	Repos map[int]*domain.User
}

func (r UserRepoMap) FindByUUID(id int) (*domain.User, error) {
	if v, ok := r.Repos[id]; ok {
		return v, nil
	}
	return nil, errors.New("user not found")
}

package repo

import (
	"errors"
	"http_auth/internal/domain"
)

type UserRepoMap struct {
	// Чисто perfomance изменение, название поставил по-спецефичнее для конкретно этого репозитория
	UserData map[int]*domain.User
}

func NewUserRepoMap(data map[int]*domain.User) *UserRepoMap {
	return &UserRepoMap{UserData: data}
}

func (r UserRepoMap) FindByUUID(id int) (*domain.User, error) {
	if v, ok := r.UserData[id]; ok {
		return v, nil
	}
	return nil, errors.New("user not found")
}

type OtherUserRepo struct{}

func NewOtherUserRepo() *OtherUserRepo {
	return &OtherUserRepo{}
}

func (r *OtherUserRepo) FindByUUID(id int) (*domain.User, error) {
	return &domain.User{ID: id, Name: "Drugoi", Surname: "Repo"}, nil
}

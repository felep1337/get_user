package repo

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"http_auth/internal/domain"
	"log"
)

type UserRepoMap struct {
	// Чисто perfomance изменение, название поставил по-спецефичнее для конкретно этого репозитория
	UserData map[int]*domain.User //Подключение к постгресу
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

type UserRepoPostgres struct {
	UserData *sql.DB
}

func NewUserRepoPostgres(connStr string) *UserRepoPostgres {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			surname TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Println("Ошибка при создании таблицы:", err)
		return nil
	}
	_, err = db.Exec(`INSERT INTO users (name, surname) VALUES ($1, $2)`, "first", "user")
	if err != nil {
		log.Println("Ошибка при вставке:", err)
		return nil
	}
	return &UserRepoPostgres{UserData: db}
}

func (r *UserRepoPostgres) FindByUUID(id int) (*domain.User, error) {

	rows, err := r.UserData.Query("SELECT name, surname FROM users WHERE id=$1", id)
	if err != nil {
		fmt.Println(err, "1")
		return nil, err
	}
	defer rows.Close()
	var (
		Name    string
		Surname string
	)
	for rows.Next() {
		err := rows.Scan(&Name, &Surname)
		if err != nil {
			fmt.Println(err, "2")
			return nil, err
		}
		fmt.Println(id, Name, Surname)
	}
	return &domain.User{ID: id, Name: Name, Surname: Surname}, nil
}

type OtherUserRepo struct{}

func NewOtherUserRepo() *OtherUserRepo {
	return &OtherUserRepo{}
}

func (r *OtherUserRepo) FindByUUID(id int) (*domain.User, error) {
	return &domain.User{ID: id, Name: "Drugoi", Surname: "Repo"}, nil
}

package main

import (
	"fmt"
	"http_auth/internal/domain"
	"http_auth/internal/handler/http_handler"
	"http_auth/internal/repo"
	"http_auth/internal/use_case"
	"net/http"
)

func main() {

	DB := map[int]*domain.User{
		1: {ID: 1, Name: "Alice", Surname: "Dvachevckaya"},
		2: {ID: 2, Name: "Bob", Surname: "Marley"},
	}
	Repository := &repo.UserRepoMap{Repos: DB}
	UseCase := &use_case.UserRepo{Data: Repository}
	handler := &http_handler.Handler{Data: UseCase}
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", handler)

}

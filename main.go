package main

import (
	"fmt"
	"http_auth/internal/handler/http_handler"
	"http_auth/internal/repo"
	"http_auth/internal/use_case"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//DB := map[int]*domain.User{
	//	1: {ID: 1, Name: "Alice", Surname: "Dvachevckaya"},
	//	2: {ID: 2, Name: "Bob", Surname: "Marley"},
	//	3: {ID: 3, Name: "Steve", Surname: "Huis"},
	//}
	// локальные переменные лучше именовать с маленькой буквы camelCase-ом

	/*
		Вот, следи за руками - отказавшись от зависимости от конкретной реализации,
		прокинув в useCase другой репозиторий - мы не меняя useCase, не меняя handler
		полностью поменяли функционал получения данных - в этом и сила инверсии зависимостей - модульность.
		Это позволяет собирать твои проекты из чего угодно, просто реализовывая контракт, допустим ты можешь на лету
		поменять БД с clickhouse на PostgreSQL и тд и тп. Работает это на все остальные уровни тоже
	*/
	//repository := repo.NewUserRepoMap(DB)
	connStr := "user=postgres password=secret dbname=mydb host=localhost port=5432 sslmode=disable"

	DBRepository := repo.NewUserRepoPostgres(connStr)
	useCase := use_case.NewUserRepo(DBRepository)
	handler := http_handler.NewHandler(useCase)
	fmt.Println("Server is running at http://localhost:8080")

	// Вот так обычно делают handler-ы в проде.
	http.HandleFunc("/", handler.FindUserByID)
	return http.ListenAndServe(":8080", nil)
}

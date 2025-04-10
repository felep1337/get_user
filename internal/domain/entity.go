package domain

type User struct {
	ID      int
	Name    string
	Surname string
}

type FindByUUID interface {
	FindByUUID(int) (*User, error)
}

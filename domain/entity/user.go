package entity

type User struct {
	ID       string
	Username string
	Email    string
	Password string
	Roles    []string
}

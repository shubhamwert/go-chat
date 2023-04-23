package chatApp

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int
}

var UsersList map[string]User

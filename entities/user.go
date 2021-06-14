package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (it *User) HidePassword() *User {
	u := *it
	u.Password = ""
	return &u
}

package model

type User struct {
	Username string
	Name     string
	Password string
}

func (u User) GetKey() string {
	return u.Username
}

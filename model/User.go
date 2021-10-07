package model

type User struct {
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

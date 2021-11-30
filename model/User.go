package model

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Words    string `json:"words"`
}

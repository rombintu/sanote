package models

type User struct {
	Model
	Login    string `json:"login"`
	Password string `json:"password"`
}

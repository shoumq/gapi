package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Admin    bool   `json:"is_admin"`
	Password string `json:"-"`
}

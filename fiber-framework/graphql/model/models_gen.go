// Code generated by github.com/arsmn/fastgql, DO NOT EDIT.

package model

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

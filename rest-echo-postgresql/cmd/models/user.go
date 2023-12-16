package models

import "time"

type IUser interface {
	SetID(id int)
	GetName() string
	GetEmail() string
	GetPassword() string
}

func NewUser(
	name string ,
	email string ,
	password string,
) IUser {
	return &user {
		Name: name,
		Email: email,
		Password: password,
	}
}

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *user) GetName() string {
	return u.Name
}

func (u *user) GetEmail() string {
	return u.Email
}

func (u *user) GetPassword() string {
	return u.Password
}

func (u *user) SetID(id int) {
	u.ID = id
}
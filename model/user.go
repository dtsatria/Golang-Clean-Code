package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Divisi    string    `json:"divisi" binding:"required"`
	Jabatan   string    `json:"jabatan" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password,omitempty" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "employee" || u.Role == "GA"
}

func (u User) IsEmpty() bool {
	return u.Name == "" || u.Divisi == "" || u.Jabatan == "" || u.Email == "" || u.Password == ""
}

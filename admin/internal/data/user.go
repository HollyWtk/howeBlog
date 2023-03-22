package data

import "time"

type User struct {
	Id         int64
	Username   string
	Password   string
	Nickname   string
	Avatar     string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
	Role       string
}

func (u *User) TableName() string {
	return "user"
}

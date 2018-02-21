package model

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

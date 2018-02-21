package row

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

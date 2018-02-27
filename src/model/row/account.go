package row

import "time"

type Account struct {
	Id          int       `xorm:"not null pk autoincr INT(10)"  json:"id"`
	Name        string    `xorm:"not null VARCHAR(100)" validate:"required" json:"name"`
	Description string    `xorm:"VARCHAR(100)" json:"description"`
	Text        string    `xorm:"VARCHAR(100)" json:"text"`
	CreatedAt   time.Time `xorm:"<-" json:"created_at"`
	UpdatedAt   time.Time `xorm:"<-" json:"updated_at"`
}

type Accounts = []Account

package row

import "time"

type Account struct {
	Id          int
	Name        string
	Description string
	Text        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Accounts = []Account

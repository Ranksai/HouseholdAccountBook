package row

import "time"

type Item struct {
	Id          int
	Name        string
	Description string
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Items = []Item

package row

import "time"

type AccountItem struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	AccountId int
	ItemId    int
}

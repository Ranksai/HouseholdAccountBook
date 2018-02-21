package row

import "time"

type AccountAccountType struct {
	Id            int
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AccountId     int
	AccountTypeId int
}

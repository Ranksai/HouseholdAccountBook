package row

import "time"

type HouseholdAccountBookAccount struct {
	Id                     int
	Name                   string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	HouseholdAccountBookId int
	AccountId              int
}

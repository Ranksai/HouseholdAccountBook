package row

import "time"

// TODO: Enum
type AccountType struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

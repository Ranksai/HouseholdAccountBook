package entity

import "github.com/Ranksai/HouseholdAccountBook/src/model/row"

type HouseholdAccountBook struct {
	row.HouseholdAccountBook
	Accounts row.Accounts
}

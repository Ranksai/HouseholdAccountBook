package entity

import "github.com/Ranksai/HouseholdAccountBook/src/model/row"

type Account struct {
	row.Account
	Items       row.Items
	Amount      int
	AccountType row.AccountType
}

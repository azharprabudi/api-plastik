package model

// ExpenseImageCreate ...
type ExpenseImageCreate struct {
	Value string `db:"value"`
}

// ExpenseImageRead ...
type ExpenseImageRead struct {
	ID    int    `db:"key"`
	Value string `db:"value"`
}

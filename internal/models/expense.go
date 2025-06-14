package models

type Expense struct {
	ID            int64   `json:"id"`
	Expense       string  `json:"expense"`
	ExpenseAmount float64 `json:"expense_amount"`
	DueDate       int     `json:"due_date"`
}

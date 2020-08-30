package schemas

// FinanceAccountTypeCreate ...
type FinanceAccountTypeCreate struct {
	Name        string `validate:"required"`
	Description string `validate:"omitempty"`
}

// FinanceAccountTypeUpdate ...
type FinanceAccountTypeUpdate struct {
	Name        string `validate:"omitempty"`
	Description string `validate:"omitempty"`
}

// TypeID is param uri validation for id
type TypeID struct {
	ID uint `uri:"id" binding:"required"`
}

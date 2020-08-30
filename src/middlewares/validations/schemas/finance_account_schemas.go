package schemas

// FinanceAccountCreate ...
type FinanceAccountCreate struct {
	Name        string `validate:"required"`
	Description string `validate:"omitempty"`
	TypeID      uint   `validate:"required,gt=0"`
}

// FinanceAccountUpdate ...
type FinanceAccountUpdate struct {
	Name        string `validate:"omitempty"`
	Description string `validate:"omitempty"`
	TypeID      uint   `validate:"omitempty,gt=0"`
}

// ID is param uri validation for id
type ID struct {
	ID uint `uri:"id" binding:"required"`
}

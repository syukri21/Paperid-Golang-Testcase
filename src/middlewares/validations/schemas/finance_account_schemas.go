package schemas

// FinanceAccountCreate ...
type FinanceAccountCreate struct {
	Name        string `validate:"required"`
	Description string `validate:"omitempty"`
	TypeID      uint   `validate:"required,gt=0" form:"typeId" json:"typeId"`
}

// FinanceAccountUpdate ...
type FinanceAccountUpdate struct {
	Name        string `validate:"omitempty"`
	Description string `validate:"omitempty"`
	TypeID      uint   `validate:"omitempty,gt=0" form:"typeId" json:"typeId"`
}

// ID is param uri validation for id
type ID struct {
	ID string `uri:"id" binding:"required"`
}

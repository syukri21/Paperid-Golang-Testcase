package schemas

// Pagination ...
type Pagination struct {
	Offset int `validate:"omitempty,numeric" form:"offset"`
	Limit  int `validate:"omitempty,numeric" form:"limit"`
}

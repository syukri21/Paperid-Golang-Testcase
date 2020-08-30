package schemas

// Pagination ...
type Pagination struct {
	Offset int `validate:"omitempty,numeric" json:"default=0"`
	Limit  int `validate:"omitempty,numeric" json:"default=2"`
}

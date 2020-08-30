package schemas

// ProfileUpdate ...
type ProfileUpdate struct {
	FirstName   string `validate:"required"`
	LastName    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Address     string `validate:"required"`
}

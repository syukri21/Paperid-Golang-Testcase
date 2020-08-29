package schemas

// Signup ...
type Signup struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

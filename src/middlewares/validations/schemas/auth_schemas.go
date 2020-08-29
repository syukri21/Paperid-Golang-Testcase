package schemas

// Signup ...
type Signup struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

// Signin ...
type Signin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

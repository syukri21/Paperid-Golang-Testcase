package flag

// AuthFlag is response code and message for user controller
type AuthFlag struct {
	Message string
	Flag    string
	Error
}

var (
	// SigninSuccess is response code and message
	SigninSuccess = AuthFlag{
		Message: "Success signin",
	}
	// SignupSuccess is response code and message
	SignupSuccess = AuthFlag{
		Message: "Success signup",
	}
	// SignoutSuccess is response code and message
	SignoutSuccess = AuthFlag{
		Message: "Success signout",
	}

	// SigninErrors is response code and message
	SigninErrors = AuthFlag{
		Error: Error{
			Message: "Something went wrong",
			Flag:    "BAD_REQUEST",
		},
	}
	// SignupErrors is response code and message
	SignupErrors = AuthFlag{
		Error: Error{
			Message: "Something went wrong",
			Flag:    "BAD_REQUEST",
		},
	}
	// SignoutErrors is response code and message
	SignoutErrors = AuthFlag{
		Error: Error{
			Message: "Something went wrong",
			Flag:    "BAD_REQUEST",
		},
	}
)

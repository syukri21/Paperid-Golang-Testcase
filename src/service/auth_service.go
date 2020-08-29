package service

import (
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/validations/schemas"
)

// AuthService -> AuthService
type AuthService struct {
	UserRepository repositories.UserRepository
}

// AuthServiceInstance -> AuthServiceInstance
func AuthServiceInstance() AuthService {
	return AuthService{UserRepository: repositories.UserRepositoryInstance()}
}

// Signup -> Signup
func (s *AuthService) Signup(user entity.User) repositories.GetUser {
	ID := &user.ID
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{
		Email: user.Email,
		ID:    ID,
	})

	if userExist {
		exception.Conflict("User conflict", []map[string]interface{}{
			{"message": "User with this email already exist", "flag": "USER_ALREADY_EXIST"},
		})
	}
	newUser := s.UserRepository.CreateUser(user)

	return newUser
}

// Signin ...
func (s *AuthService) Signin(params schemas.Signin) repositories.GetUser {

	user, err := s.UserRepository.GetUserByEmail(params.Email)
	if err != nil {
		exception.NotFound("User or Password Wrong", []map[string]interface{}{
			{"message": "user or password wrong", "flag": "USER_OR_PASSWORD_WRONG"},
		})
	}

	match := user.ComparePassword(params.Password)
	if !match {
		exception.NotFound("User or Password Wrong", []map[string]interface{}{
			{"message": "user or password wrong", "flag": "USER_OR_PASSWORD_WRONG"},
		})
	}

	token, err := user.GenerateJWT()
	err = s.UserRepository.Conn.Model(&user).Update("jwt_token", token).Error
	if err != nil {
		exception.NotFound("Something Went Wrong", []map[string]interface{}{
			{"message": "Something Went Wrong", "flag": "SOMETHING_WRONG"},
		})
	}

	return repositories.GetUser{
		Email:     user.Email,
		ID:        user.ID,
		JwtToken:  *user.JwtToken,
		DeletedAt: user.DeletedAt,
	}
}

// Signout ...
func (s *AuthService) Signout(user entity.User) bool {
	err := s.UserRepository.Conn.Model(&user).Update("jwt_token", "").Error
	return err == nil
}

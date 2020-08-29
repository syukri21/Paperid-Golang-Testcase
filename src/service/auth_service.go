package service

import (
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
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
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{
		Email: user.Email,
		ID:    user.ID,
	})

	if (userExist != entity.User{}) {
		exception.Conflict("User conflict", []map[string]interface{}{
			{"message": "User with this email already exist", "flag": "USER_ALREADY_EXIST"},
		})
	}

	newUser := s.UserRepository.CreateUser(user)
	return newUser
}

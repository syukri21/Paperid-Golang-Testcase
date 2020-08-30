package service

import (
	uuid "github.com/satori/go.uuid"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
)

// UserService ...
type UserService struct {
	UserRepository repositories.UserRepository
}

// UserServiceInstance -> UserServiceInstance
func UserServiceInstance() UserService {
	return UserService{UserRepository: repositories.UserRepositoryInstance()}
}

// GetOne ...
func (r *UserService) GetOne(id uuid.UUID) (result interface{}) {
	result, err := r.UserRepository.GetUserProfileByID(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

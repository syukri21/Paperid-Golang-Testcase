package service

import (
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
)

// ProfileService ...
type ProfileService struct {
	ProfileRepository repositories.ProfileRepository
}

// ProfileServiceInstance ...
func ProfileServiceInstance() ProfileService {
	return ProfileService{
		ProfileRepository: repositories.ProfileRepositoryInstance(),
	}
}

// UpdateOrCreate ...
func (s *ProfileService) UpdateOrCreate(data entity.Profile) map[string]interface{} {

	isExist := s.ProfileRepository.Exist(data.UserID)

	var err error

	if isExist {
		_, err = s.ProfileRepository.Update(data.UserID, data)
		if err != nil {
			exception.BadRequest("Something Went Wrong", []map[string]interface{}{
				{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
			})
		}
	} else {
		_, err = s.ProfileRepository.Create(data)
		if err != nil {
			exception.BadRequest("Something Went Wrong", []map[string]interface{}{
				{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
			})
		}
	}

	return map[string]interface{}{"message": "success"}
}

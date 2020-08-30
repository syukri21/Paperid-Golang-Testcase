package service

import (
	"github.com/satori/go.uuid"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
)

// FinanceAccountService ...
type FinanceAccountService struct {
	FinanceAccountRepository repositories.FinanceAccountRepository
}

// FinanceAccountServiceInstance ...
func FinanceAccountServiceInstance() FinanceAccountService {
	return FinanceAccountService{
		FinanceAccountRepository: repositories.FinanceAccountRepositoryInstance(),
	}
}


// GetAll ...
func (s *FinanceAccountService) GetAll() (results []entity.FinanceAccount) {
	results, err := s.FinanceAccountRepository.GetAll()
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// GetByID ...
func (s *FinanceAccountService) GetByID(id uuid.UUID) (result entity.FinanceAccount) {
	result, err := s.FinanceAccountRepository.GetByID(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// Create ...
func (s *FinanceAccountService) Create(data entity.FinanceAccount) (result entity.FinanceAccount) {
	result, err := s.FinanceAccountRepository.Create(data)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// Update ...
func (s *FinanceAccountService) Update(id uuid.UUID, data entity.FinanceAccount) map[string]interface{} {

	isExist := s.FinanceAccountRepository.Exist(id)

	if !isExist {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": "Type not exist or already delete", "flag": "BAD_REQUEST"},
		})
	}

	success, _ := s.FinanceAccountRepository.Update(id, data)
	if !success {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return map[string]interface{}{"message": "success"}
}

// Delete ...
func (s *FinanceAccountService) Delete(id uuid.UUID) map[string]interface{} {

	isExist := s.FinanceAccountRepository.Exist(id)

	if !isExist {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": "Type not exist or already delete", "flag": "BAD_REQUEST"},
		})
	}

	success, _ := s.FinanceAccountRepository.DeleteByID(id)
	if !success {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	return map[string]interface{}{"message": "success"}
}

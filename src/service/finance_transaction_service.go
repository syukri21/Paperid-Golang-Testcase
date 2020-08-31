package service

import (
	"github.com/satori/go.uuid"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
)

// FinanceTransactionService ...
type FinanceTransactionService struct {
	FinanceTransactionRepository repositories.FinanceTransactionRepository
}

// FinanceTransactionServiceInstance ...
func FinanceTransactionServiceInstance() FinanceTransactionService {
	return FinanceTransactionService{
		FinanceTransactionRepository: repositories.FinanceTransactionRepositoryInstance(),
	}
}

// GetAll ...
func (s *FinanceTransactionService) GetAll(p schemas.Pagination, date schemas.Summary) (results []entity.FinanceTransaction) {
	results, err := s.FinanceTransactionRepository.GetAll(p, date)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// GetByID ...
func (s *FinanceTransactionService) GetByID(id string) (result entity.FinanceTransaction) {
	u, err := uuid.FromString(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	result, err = s.FinanceTransactionRepository.GetByID(u)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// Create ...
func (s *FinanceTransactionService) Create(data entity.FinanceTransaction) (result entity.FinanceTransaction) {
	result, err := s.FinanceTransactionRepository.Create(data)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return
}

// Update ...
func (s *FinanceTransactionService) Update(id string, data entity.FinanceTransaction) map[string]interface{} {

	u, err := uuid.FromString(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	isExist := s.FinanceTransactionRepository.Exist(u)

	if !isExist {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": "Type not exist or already delete", "flag": "BAD_REQUEST"},
		})
	}

	success, _ := s.FinanceTransactionRepository.Update(u, data)
	if !success {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return map[string]interface{}{"message": "success"}
}

// Delete ...
func (s *FinanceTransactionService) Delete(id string) map[string]interface{} {
	u, err := uuid.FromString(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	isExist := s.FinanceTransactionRepository.Exist(u)

	if !isExist {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": "Type not exist or already delete", "flag": "BAD_REQUEST"},
		})
	}

	success, _ := s.FinanceTransactionRepository.DeleteByID(u)
	if !success {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	return map[string]interface{}{"message": "success"}
}

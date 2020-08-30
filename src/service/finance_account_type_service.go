package service

import (
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
)

// FinanceAccountTypeService -> FinanceAccountTypeService
type FinanceAccountTypeService struct {
	FinanceAccountTypeRepository repositories.FinanceAccountTypeRepository
}

// FinanceAccountTypeServiceInstance -> FinanceAccountTypeServiceInstance
func FinanceAccountTypeServiceInstance() FinanceAccountTypeService {
	return FinanceAccountTypeService{
		FinanceAccountTypeRepository: repositories.FinanceAccountTypeRepositoryInstance(),
	}
}

// GetAll ...
func (s *FinanceAccountTypeService) GetAll() []entity.FinanceAccountType {
	types := s.FinanceAccountTypeRepository.GetTypes()
	return types
}

// Create ...
func (s *FinanceAccountTypeService) Create(Type schemas.FinanceAccountTypeCreate) entity.FinanceAccountType {
	types, err := s.FinanceAccountTypeRepository.Create(entity.FinanceAccountType{
		Name:        Type.Name,
		Description: Type.Description,
	})
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return types
}

// GetByID ...
func (s *FinanceAccountTypeService) GetByID(id uint) entity.FinanceAccountType {
	types, err := s.FinanceAccountTypeRepository.GetTypeByID(id)
	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return types
}

// Update ...
func (s *FinanceAccountTypeService) Update(id uint, t entity.FinanceAccountType) map[string]interface{} {

	isExist := s.FinanceAccountTypeRepository.Exist(id)

	if isExist {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": "Type Already Exist", "flag": "BAD_REQUEST"},
		})
	}

	_, err := s.FinanceAccountTypeRepository.Update(id, repositories.FinanceUpdataParam{
		Name:        &t.Name,
		Description: &t.Description,
	})

	if err != nil {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}
	return map[string]interface{}{"message": "success"}
}

// Delete ...
func (s *FinanceAccountTypeService) Delete(id uint) map[string]interface{} {

	success, _ := s.FinanceAccountTypeRepository.DeleteType(id)

	if !success {
		exception.BadRequest("Something Went Wrong", []map[string]interface{}{
			{"message": flags.DefaultError.Message, "flag": flags.DefaultError.Flag},
		})
	}

	return map[string]interface{}{"message": "success"}
}

package service

import (
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"
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
func (s *FinanceAccountTypeService) Create(Type entity.FinanceAccountType) []entity.FinanceAccountType {
	types := s.FinanceAccountTypeRepository.()
	return types
}

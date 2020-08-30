package service

import "github.com/syukri21/Paperid-Golang-Testcase/src/repositories"

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

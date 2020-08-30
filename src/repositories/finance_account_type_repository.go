package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"

	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
)

// FinanceAccountTypeRepository -> the propose of user repository is handling query for user entity
type FinanceAccountTypeRepository struct {
	Conn *gorm.DB
}

// FinanceAccountTypeRepositoryInstance ->
func FinanceAccountTypeRepositoryInstance() FinanceAccountTypeRepository {
	conn := db.GetDB().Table("finance_account_types")
	return FinanceAccountTypeRepository{Conn: conn}
}

// GetTypes ->
func (r *FinanceAccountTypeRepository) GetTypes() []entity.FinanceAccountType {
	financeAccountTypes := []entity.FinanceAccountType{}
	r.Conn.Find(&financeAccountTypes)
	return financeAccountTypes
}

// GetTypeByID ->
func (r *FinanceAccountTypeRepository) GetTypeByID(id int) entity.FinanceAccountType {
	financeAccountType := entity.FinanceAccountType{}
	r.Conn.Where("id = ?", id).First(&financeAccountType)
	return financeAccountType
}

// UpdateType ->
func (r *FinanceAccountTypeRepository) UpdateType(Type *entity.FinanceAccountType) (entity.FinanceAccountType, error) {
	financeAccountType := entity.FinanceAccountType{}
	err := r.Conn.Update(Type).Error
	return financeAccountType, err
}

// DeleteType ->
func (r *FinanceAccountTypeRepository) DeleteType(id int) (success bool) {
	time := time.Now()
	err := r.Conn.Where("id = ?", id).Update(&entity.FinanceAccountType{
		DeletedAt: &time,
	}).Error
	return err == nil
}

// Create ->
func (r *FinanceAccountTypeRepository) Create(Type entity.FinanceAccountType) (entity.FinanceAccountType, error) {
	err := r.Conn.Create(&Type).Error
	return Type, err
}

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
	r.Conn.Where("deleted_at IS NULL").Find(&financeAccountTypes)
	return financeAccountTypes
}

// GetTypeByID ->
func (r *FinanceAccountTypeRepository) GetTypeByID(id uint) (financeAccountType entity.FinanceAccountType, err error) {
	err = r.Conn.Where("id = ? AND deleted_at IS NULL", id).First(&financeAccountType).Error
	return
}

// Update ->
func (r *FinanceAccountTypeRepository) Update(id uint, params entity.FinanceAccountType) (bool, error) {
	time := time.Now()
	params.UpdatedAt = time
	err := r.Conn.Where("id = ?", id).Update(params).Error
	return err == nil, err
}

// DeleteType ->
func (r *FinanceAccountTypeRepository) DeleteType(id uint) (bool, error) {
	time := time.Now()
	err := r.Conn.Where("id = ?", id).Update(&entity.FinanceAccountType{
		DeletedAt: &time,
	}).Error
	return err == nil, err
}

// Create ->
func (r *FinanceAccountTypeRepository) Create(Type entity.FinanceAccountType) (entity.FinanceAccountType, error) {
	err := r.Conn.Create(&Type).Error
	return Type, err
}

// Exist ...
func (r *FinanceAccountTypeRepository) Exist(id uint) (exist bool) {

	t := entity.FinanceAccountType{}
	if err := r.Conn.Select("id").Where("id = ?", id).First(&t).Error; err != nil {
		exist = false
	}

	if t.ID > 0 {
		exist = true
	} else {
		exist = false
	}

	return

}

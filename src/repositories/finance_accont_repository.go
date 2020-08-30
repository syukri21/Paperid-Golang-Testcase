package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// FinanceAccountRepository ...
type FinanceAccountRepository struct {
	Conn *gorm.DB
}

// FinanceAccountRepositoryInstance ...
func FinanceAccountRepositoryInstance() FinanceAccountRepository {
	return FinanceAccountRepository{
		Conn: db.GetDB().Table("finance_account"),
	}
}

// GetAll ->
func (r *FinanceAccountRepository) GetAll() (data []entity.FinanceAccount, err error) {
	err = r.Conn.Where("deleted_at IS NULL").Find(&data).Error
	return
}

// GetByID ->
func (r *FinanceAccountRepository) GetByID(id uuid.UUID) (data entity.FinanceAccount, err error) {
	err = r.Conn.Where("id = ? AND deleted_at IS NULL", id).First(&data).Error
	return
}

// Update ->
func (r *FinanceAccountRepository) Update(id uuid.UUID, params entity.FinanceAccount) (bool, error) {
	time := time.Now()
	params.UpdatedAt = time
	err := r.Conn.Where("id = ?", id).Update(params).Error
	return err == nil, err
}

// DeleteByID ->
func (r *FinanceAccountRepository) DeleteByID(id uuid.UUID) (bool, error) {
	time := time.Now()
	err := r.Conn.Where("id = ?", id).Update(&entity.FinanceAccount{
		DeletedAt: &time,
	}).Error
	return err == nil, err
}

// Create ->
func (r *FinanceAccountRepository) Create(params entity.FinanceAccount) (data entity.FinanceAccount, err error) {
	err = r.Conn.Create(&params).Error
	data = params
	return
}

// Exist ...
func (r *FinanceAccountRepository) Exist(id uuid.UUID) (exist bool) {

	data := entity.FinanceAccount{}
	if err := r.Conn.Select("id").Where("id = ?", id).First(&data).Error; err != nil {
		exist = false
	}

	if len(data.ID) > 0 {
		exist = true
	} else {
		exist = false
	}

	return

}

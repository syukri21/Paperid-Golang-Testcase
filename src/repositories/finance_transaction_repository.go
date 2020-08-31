package repositories

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
)

// FinanceTransactionRepository ...
type FinanceTransactionRepository struct {
	Conn *gorm.DB
}

// FinanceTransactionRepositoryInstance ...
func FinanceTransactionRepositoryInstance() FinanceTransactionRepository {
	return FinanceTransactionRepository{
		Conn: db.GetDB().Table("finance_transactions"),
	}
}

// GetAll ->
func (r *FinanceTransactionRepository) GetAll(p schemas.Pagination, date schemas.Summary) (data []entity.FinanceTransaction, err error) {
	query := r.Conn.Debug().Offset(p.Offset).Limit(p.Limit)

	fmt.Printf("%s", date)
	if date.Day > 0 || date.Month > 0 {
		now := time.Now()
		after := now.AddDate(0, -date.Month, 0)
		after = after.AddDate(0, 0, -date.Day)
		query = query.Where("created_at > ?", after)
	} else {
		query = query.Where("deleted_at IS NULL")
	}

	err = query.Find(&data).Error
	return
}

// GetByID ->
func (r *FinanceTransactionRepository) GetByID(id uuid.UUID) (data entity.FinanceTransaction, err error) {
	err = r.Conn.Debug().Where("id = ?", id).Preload("User").Preload("FinanceAccount").First(&data).Error
	return
}

// Update ->
func (r *FinanceTransactionRepository) Update(id uuid.UUID, params entity.FinanceTransaction) (bool, error) {
	time := time.Now()
	params.UpdatedAt = time
	err := r.Conn.Where("id = ?", id).Update(params).Error
	return err == nil, err
}

// DeleteByID ->
func (r *FinanceTransactionRepository) DeleteByID(id uuid.UUID) (bool, error) {
	time := time.Now()
	err := r.Conn.Where("id = ?", id).Update(&entity.FinanceTransaction{
		DeletedAt: &time,
	}).Error
	return err == nil, err
}

// Create ->
func (r *FinanceTransactionRepository) Create(params entity.FinanceTransaction) (data entity.FinanceTransaction, err error) {
	err = r.Conn.Create(&params).Error
	data = params
	return
}

// Exist ...
func (r *FinanceTransactionRepository) Exist(id uuid.UUID) (exist bool) {

	data := entity.FinanceTransaction{}
	if err := r.Conn.Select("id").Where("id = ?", id).First(&data).Error; err != nil {
		exist = false
	}

	exist = len(data.ID) > 0
	return
}

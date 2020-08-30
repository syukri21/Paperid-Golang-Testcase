package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
)

// ProfileRepository ...
type ProfileRepository struct {
	Conn *gorm.DB
}

// ProfileRepositoryInstance ...
func ProfileRepositoryInstance() ProfileRepository {
	return ProfileRepository{
		Conn: db.GetDB().Table("profiles"),
	}
}

// Create ...
func (r *ProfileRepository) Create(data entity.Profile) (result entity.Profile, err error) {
	err = r.Conn.Create(&data).Error
	result = data
	return
}

// Update ...
func (r *ProfileRepository) Update(id uuid.UUID, data entity.Profile) (result entity.Profile, err error) {
	time := time.Now()
	data.UpdatedAt = time
	err = r.Conn.Where("id = ? OR user_id = ?", id, id).Update(&data).Error
	return data, err
}

// Exist ...
func (r *ProfileRepository) Exist(id uuid.UUID) (exist bool) {
	data := entity.FinanceTransaction{}
	rowsAffected := r.Conn.Debug().Select("id").Where("id = ? OR user_id = ?", id, id).First(&data).RowsAffected
	return rowsAffected > 0
}

package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Profile -> profile entity schema
type Profile struct {
	ID          uuid.UUID  `gorm:"primary_key;type:char(36)" json:"id"`
	UserID      uuid.UUID  `gorm:"unique_key;type:char(36)" json:"user_id"`
	FirstName   string     `gorm:"type:varchar(255)" json:"firstName" faker:"first_name"`
	LastName    string     `gorm:"type:varchar(255)" json:"lastName" faker:"last_name"`
	PhoneNumber string     `gorm:"type:varchar(15)" json:"phone_number" faker:"last_name"`
	Address     string     `gorm:"type:text" json:"address"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate ...
func (m *Profile) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.NewV4()
	return
}

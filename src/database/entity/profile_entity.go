package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

// Profile -> profile entity schema
type Profile struct {
	ID        uuid.UUID  `gorm:"primary_key;type:char(36)" json:"id"`
	FirstName string     `gorm:"type:varchar(255)" json:"firstName" faker:"first_name"`
	LastName  string     `gorm:"type:varchar(255)" json:"lastName" faker:"last_name"`
	Balance   float64    `gorm:"type:varchar(255)" json:"balance" faker:"oneof: 47463, 993747, 11131997"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

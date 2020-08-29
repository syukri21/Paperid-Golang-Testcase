package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User -> user entity schema
type User struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid" json:"id"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Password  string     `gorm:"type:varchar(255)" json:"password" faker:"password"`
	JwtToken  string     `gorm:"type:varchar(255)" json:"JwtToken" faker:"jwt"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

package entity

import "time"

// User -> user entity schema
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Password  string     `gorm:"type:varchar(255)" json:"password" faker:"password"`
	JwtToken  string     `gorm:"type:varchar(255)" json:"JwtToken" faker:"jwt"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

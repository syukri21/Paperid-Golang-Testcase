package entity

import (
	"time"
)

// FinanceAccountType -> FinanceAccountType entity schema
type FinanceAccountType struct {
	ID          uint       `gorm:"primary_key;auto_increment:true" json:"id"`
	Name        string     `gorm:"type:varchar(255)" json:"name" faker:"first_name"`
	Description string     `gorm:"type:text" json:"description" faker:"oneof: 47463, 993747, 11131997"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

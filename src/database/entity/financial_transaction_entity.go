package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// FinanceTransaction -> FinanceTransaction entity schema
type FinanceTransaction struct {
	ID               uuid.UUID      `gorm:"primary_key;type:char(36)" json:"id"`
	FinanceAccountID uuid.UUID      `gorm:"type:char(36)" json:"finance_account_id"`
	UserID           uuid.UUID      `gorm:"type:char(36)" json:"user_id"`
	Name             string         `gorm:"type:varchar(255)" json:"name" faker:"first_name"`
	Amount           float64        `gorm:"type:bigint" json:"amount"`
	Description      string         `gorm:"type:text" json:"description"`
	CreatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time     `json:"deleted_at,omitempty"`
	FinanceAccount   FinanceAccount `json:"finance_account,omitempty"`
	User             User           `json:"user,omitempty"`
}

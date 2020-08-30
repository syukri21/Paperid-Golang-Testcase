package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// FinanceTransaction -> FinanceTransaction entity schema
type FinanceTransaction struct {
	ID               uuid.UUID               `gorm:"primary_key;type:char(36)" json:"id"`
	FinanceAccountID uuid.UUID               `gorm:"type:char(36)" json:"finance_account_id"`
	UserID           uuid.UUID               `gorm:"type:char(36)" json:"user_id"`
	Name             string                  `gorm:"type:varchar(255)" json:"name" faker:"first_name"`
	Amount           float64                 `gorm:"type:bigint" json:"amount"`
	Description      string                  `gorm:"type:text" json:"description"`
	CreatedAt        time.Time               `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time               `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time              `json:"deleted_at,omitempty"`
	FinanceAccount   *FinanceAccount         `json:"finance_account,omitempty"`
	User             *FinanceTransactionUser `json:"user,omitempty"`
}

// FinanceTransactionUser ...
type FinanceTransactionUser struct {
	ID              uuid.UUID        `gorm:"primary_key;type:char(36);" json:"id"`
	Email           string           `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	FinanceAccounts []FinanceAccount `json:"finance_accounts,omitempty"`
	Profile         *Profile         `json:"profile,omitempty"`
}

// TableName ...
func (FinanceTransactionUser) TableName() string {
	return "users"
}

// BeforeCreate ...
func (m *FinanceTransaction) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.NewV4()
	return
}

package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// FinanceAccount -> FinanceAccount entity schema
type FinanceAccount struct {
	ID                  uuid.UUID            `gorm:"primary_key;type:char(36)" json:"id"`
	UserID              uuid.UUID            `gorm:"type:char(36)" json:"user_id"`
	Name                string               `gorm:"type:varchar(255)" json:"name" faker:"first_name"`
	TypeID              uint                 `gorm:"type:int" json:"type_id" faker:"oneof: 1, 2"`
	Description         string               `gorm:"type:text" json:"description" faker:"oneof: 47463, 993747, 11131997"`
	CreatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt           *time.Time           `json:"deleted_at,omitempty"`
	Type                FinanceAccountType   `json:"type,omitempty"`
	User                User                 `json:"user,omitempty"`
	FinanceTransactions []FinanceTransaction `json:"finance_transactions,omitempty"`
}

// BeforeCreate ...
func (m *FinanceAccount) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.NewV4()
	return
}

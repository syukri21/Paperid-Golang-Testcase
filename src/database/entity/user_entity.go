package entity

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"

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

func (m *User) hashingPassword() (ok bool) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 4)
	if err != nil {
		return false
	}
	m.Password = string(bytes)
	return true
}

// BeforeCreate ...
func (m *User) BeforeCreate(ctx *gorm.DB) (err error) {
	ok := m.hashingPassword()
	if !ok {
		err = errors.New("can't save invalid data")
	}
	return err
}

// ComparePassword ...
func (m *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

package entity

import (
	"time"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

// User -> user entity schema
type User struct {
	ID        uuid.UUID  `gorm:"primary_key;type:char(36);" json:"id"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Password  string     `gorm:"type:varchar(255)" json:"password" faker:"password"`
	JwtToken  *string    `gorm:"type:varchar(255)" json:"JwtToken" faker:"jwt"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (m *User) uuidGenerateV4() uuid.UUID {
	uuid := uuid.NewV4()
	return uuid
}

func (m *User) hashingPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 4)
	if err != nil {
		panic("error make password")
	}
	return string(bytes)
}

// BeforeCreate ...
func (m *User) BeforeCreate(ctx *gorm.DB) (err error) {
	m.Password = m.hashingPassword()
	m.ID = m.uuidGenerateV4()
	return
}

// ComparePassword ...
func (m *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

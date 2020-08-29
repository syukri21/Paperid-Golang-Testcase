package repositories

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"

	db "github.com/syukri21/Paperid-Golang-Testcase/src/database"
)

// UserRepository -> the propose of user repository is handling query for user entity
type UserRepository struct {
	Conn *gorm.DB
}

// GetUser -> GetUser entity schema
type GetUser struct {
	ID        uint       `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// UserRepositoryInstance -> user repository instance to get user table connection
func UserRepositoryInstance() UserRepository {
	return UserRepository{Conn: db.GetDB().Table("users")}
}

// CreateUser -> CreateUser
func (r *UserRepository) CreateUser(user entity.User) GetUser {
	r.Conn.Create(&user)
	userCreated := GetUser{}
	r.Conn.Select("id, email, deleted_at").Where("id=?").First(&userCreated)
	return userCreated
}

// UserExistParams -> Optional params for user exist
type UserExistParams struct {
	Email string
	ID    *uuid.UUID
}

// UserExist -> method to check if user already exist in database by email or id
func (r *UserRepository) UserExist(param UserExistParams) entity.User {
	user := entity.User{}
	if param.ID == nil {
		r.Conn.Select("email").Where(&entity.User{Email: param.Email}).First(&user)
	} else {
		r.Conn.Select("id").Where(&entity.User{ID: *param.ID}).First(&user)
	}
	return user
}

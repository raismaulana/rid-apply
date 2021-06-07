package repository

import (
	"github.com/raismaulana/rid-apply/golang-engineer/internal/config"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id string) entity.User
	InsertUser(user entity.User)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: config.Db,
	}
}

func (db *userRepository) GetUser(id string) entity.User {
	var user entity.User
	db.connection.Joins("SSO").Joins("Email").Find(&user, "ssos.userid = ?", id)
	return user
}

func (db *userRepository) InsertUser(user entity.User) {
	db.connection.Create(user)
}

package persistence

import (
	"errors"
	"strings"

	"github.com/raytr/sample-project-p/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u UserRepo) SaveUser(user entity.User) (interface{}, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("email already taken")
		}
		//any other db error
		return nil, errors.New("database error")
	}
	return user, nil
}

func (u UserRepo) GetUser(id uint64) (*entity.User, error) {
	user := new(entity.User)
	err := u.db.Where("id=?", id).Error
	return user, err
}

func (u UserRepo) GetUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := u.db.Table("users").Find(&users).Error
	return users, err
}

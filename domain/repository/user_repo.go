package repository

import "github.com/raytr/sample-project-p/domain/entity"

type UserRepository interface {
	SaveUser(user entity.User) (interface{}, error)
	GetUsers() ([]entity.User, error)
	GetUser(id uint64) (*entity.User, error)
}

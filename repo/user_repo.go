package repo

import "github.com/raytr/sample-project-p/domain/entity"

type UserRepo interface {
	GetUser(id uint64) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
}

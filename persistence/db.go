package persistence

import (
	"fmt"

	"github.com/raytr/sample-project-p/domain/entity"
	"github.com/raytr/sample-project-p/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

////closes the  database connection
//func (s *Repositories) Close() error {
//	return s.db.Close()
//}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{})
}

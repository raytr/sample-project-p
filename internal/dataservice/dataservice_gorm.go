package dataservice

import (
	"fmt"

	"github.com/raytr/sample-project-p/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataservice struct {
	db *gorm.DB
}

func NewDataServiceGorm(db *gorm.DB) DataService {
	ds := &dataservice{
		db: db,
	}
	return ds
}

func NewGormDB(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DbName,
		cfg.Password)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

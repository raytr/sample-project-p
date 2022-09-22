package factory

import (
	"sync"

	"github.com/raytr/sample-project-p/config"
	"github.com/raytr/sample-project-p/internal/dataservice"
	"gorm.io/gorm"
)

const (
	KeyDataExtractor = "DataExtractor"
	KeyDataService   = "DataService"
	KeyKafkaService  = "KafkaService"
	KeyController    = "Controller"
	KeyHandler       = "Handler"
	KeyGoogleStorage = "GoogleStorage"
)

type defaultFactory struct {
	cfg       *config.Config
	container map[string]interface{}
	db        *gorm.DB
	mu        sync.Mutex
}

func NewDefaultFactory(cfg *config.Config) (*defaultFactory, error) {
	db, err := dataservice.NewGormDB(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &defaultFactory{
		cfg:       cfg,
		db:        db,
		container: make(map[string]interface{}),
	}, nil
}

func (f defaultFactory) BuildDataService() dataservice.DataService {

	return dataservice.NewDataServiceGorm(f.db)
}

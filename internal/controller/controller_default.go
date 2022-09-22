package controller

import "github.com/raytr/sample-project-p/internal/dataservice"

type controller struct {
	ds dataservice.DataService
	//kafka     dataservice.KafkaService
	//extractor dataservice.DataExtractor
	//txEnabled bool
}

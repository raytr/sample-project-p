package factory

import "github.com/raytr/sample-project-p/internal/handler"

type Factory interface {
	BuildHandler() handler.Handler
}

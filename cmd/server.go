package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/raytr/sample-project-p/interfaces/handler"
	"github.com/raytr/sample-project-p/internal/factory"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Server",
		Run: func(cmd *cobra.Command, args []string) {
			//db, err := persistence.ConnectMongoDB(cfg.Database.MongoDB)
			//if err != nil {
			//	log.Printf("Failed when connect db: %v", err)
			//}
			//log.Println("Connect success")

			//Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName
			//repositories, err := persistence.NewRepositories(
			//	cfg.Database.User,
			//	cfg.Database.Password,
			//	cfg.Database.Port,
			//	cfg.Database.Host,
			//	cfg.Database.DbName,
			//)
			//if err != nil {
			//	panic(err)
			//}
			////defer repositories.Close()
			//if err = repositories.Automigrate(); err != nil {
			//	log.Println("Has error when migrate")
			//}
			//
			////applications := application.NewApplication(repositories)
			//userService := handler.NewUserHandler(repositories.User, cfg)

			f, err := factory.NewDefaultFactory(cfg)
			if err != nil {
				log.Fatal(err)
			}
			h := f.BuildHandler()

			e := echo.New()
			handler.BuildRouter(e, userService)

			e.Logger.Fatal(e.Start(":9000"))
		},
	}
)

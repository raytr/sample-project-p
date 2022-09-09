package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	m     *migrate.Migrate
	steps int

	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Execute/Rollback database migration scripts",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := sql.Open("postgres", cfg.Database.ConnectionString)
			if err != nil {
				panic(err)
			}
			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Fatalf("postgres.WithInstance Error : %s", err)
			}
			m, err := migrate.NewWithDatabaseInstance(
				fmt.Sprintf("file://%s", cfg.Database.MigrationsPath),
				"postgres", driver)
			if err != nil {
				log.Fatalf("Error creating db migrate svc: %s", err)
			}

			err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
			if err != nil {
				log.Fatalf("Error when up migration: %s", err)
			}
		},
	}
)

package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yonisaka/solid/database/seeder"
	"github.com/yonisaka/solid/internal/entity"
	"github.com/yonisaka/solid/pkg/sqlx"
	"log"
)

// newDBMigrate is a method command cli to run db migration
func (cmd *Command) newDBMigrate() *cli.Command {
	return &cli.Command{
		Name:  "db:migrate",
		Usage: "A command to run database migration",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)
			db, errConn := mysql.Connect()
			if errConn != nil {
				return fmt.Errorf("unable to connect to database: %w", errConn)
			}

			err := mysql.Migrate(db, entity.User{})
			if err != nil {
				return fmt.Errorf("cannot run auto migrate: %w", err)
			}

			log.Printf("database migration success")
			return nil
		},
	}
}

// newDBSeed is a method command cli to run db seed
func (cmd *Command) newDBSeed() *cli.Command {
	return &cli.Command{
		Name:  "db:seed",
		Usage: "A command to run database seed",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)
			db, errConn := mysql.Connect()
			if errConn != nil {
				return fmt.Errorf("unable to connect to database: %w", errConn)
			}

			seed := sqlx.NewSQLSeeder(&seeder.User{})
			err := seed.Run(db)
			if err != nil {
				return fmt.Errorf("cannot run auto seed: %w", err)
			}

			log.Printf("database seed success")
			return nil
		},
	}
}

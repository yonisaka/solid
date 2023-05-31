package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
	"github.com/yonisaka/solid/internal/service/user"
	"github.com/yonisaka/solid/pkg/sqlx"
	"log"
)

// GetListUser is a method command cli to get list user
func (cmd *Command) getListUser() *cli.Command {
	return &cli.Command{
		Name:  "user:list",
		Usage: "A command to get list user",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)
			db, err := mysql.Connect()
			if err != nil {
				return fmt.Errorf("unable to connect to database: %w", err)
			}

			userRepo := repositories.NewUserRepo(db)

			request := contract.Request{}
			response := user.NewUserList(userRepo).Serve(request)

			cmd.Print(response)

			return nil
		},
	}
}

// GetDetailUser is a method command cli to get detail user
func (cmd *Command) getDetailUser() *cli.Command {
	return &cli.Command{
		Name:  "user:detail",
		Usage: "A command to get detail user",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)
			db, err := mysql.Connect()
			if err != nil {
				return fmt.Errorf("unable to connect to database: %w", err)
			}

			userRepo := repositories.NewUserRepo(db)

			request := contract.Request{
				Data: map[string]interface{}{
					"id": 1,
				},
			}
			response := user.NewUserDetail(userRepo).Serve(request)

			cmd.Print(response)

			return nil
		},
	}
}

// CreateUser is a method command cli to create user
func (cmd *Command) createUser() *cli.Command {
	return &cli.Command{
		Name:  "user:create",
		Usage: "A command to create user",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)

			db, err := mysql.Connect()
			if err != nil {
				return fmt.Errorf("unable to connect to database: %w", err)
			}

			userRepo := repositories.NewUserRepo(db)

			request := contract.Request{
				Data: map[string]interface{}{
					"name":     "Yoni Saka",
					"email":    "yonisaka@gmail.com",
					"verified": 1,
				},
			}

			response := user.NewUserCreate(userRepo).Serve(request)

			cmd.Print(response)

			return nil
		},
	}
}

// UpdateUser is a method command cli to update user
func (cmd *Command) updateUser() *cli.Command {
	return &cli.Command{
		Name:  "user:update",
		Usage: "A command to update user",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)

			db, err := mysql.Connect()
			if err != nil {
				return fmt.Errorf("unable to connect to database: %w", err)
			}

			userRepo := repositories.NewUserRepo(db)

			request := contract.Request{
				Data: map[string]interface{}{
					"id":       1,
					"name":     "John Doe Update",
					"email":    "johndoe@mail.com",
					"verified": 1,
				},
			}

			response := user.NewUserUpdate(userRepo).Serve(request)

			cmd.Print(response)

			return nil
		},
	}
}

// DeleteUser is a method command cli to delete user
func (cmd *Command) deleteUser() *cli.Command {
	return &cli.Command{
		Name:  "user:delete",
		Usage: "A command to delete user",
		Action: func(c *cli.Context) error {
			mysql := sqlx.NewMySQLConnection(cmd.Config)

			db, err := mysql.Connect()
			if err != nil {
				return fmt.Errorf("unable to connect to database: %w", err)
			}

			userRepo := repositories.NewUserRepo(db)

			request := contract.Request{
				Data: map[string]interface{}{
					"id": 1,
				},
			}

			response := user.NewUserDelete(userRepo).Serve(request)

			cmd.Print(response)

			return nil
		},
	}
}

func (cmd *Command) Print(response contract.Response) {
	log.Println(fmt.Sprintf("Message: %s", response.Message))
	log.Println(fmt.Sprintf("Data: %v", response.Data))
}

package main

import (
	"github.com/joho/godotenv"
	"github.com/yonisaka/solid/cmd"
	"github.com/yonisaka/solid/config"
	"log"
	"os"
)

func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()
	command := cmd.NewCommand(
		cmd.WithConfig(cfg),
	)

	app := cmd.NewCLI()
	app.Commands = command.Build()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}

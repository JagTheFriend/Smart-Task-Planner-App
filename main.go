package main

import (
	"log/slog"
	"smart-task-planner/cmd/server"
	"smart-task-planner/cmd/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	utils.SetupLogger()
	slog.Info("Hello World!")
	server.StartServer()
}

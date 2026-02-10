package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"smart-task-planner/cmd/server"
)

func main() {
	fmt.Println("Hello World!")
	server.StartServer()
}

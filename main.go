package main

import (
	"devcode-todolist/core"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	core.Run()
}

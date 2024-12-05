package main

import (
	"client/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("failed to init client: %s", err.Error())
	}
}

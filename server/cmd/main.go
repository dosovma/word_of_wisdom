package main

import (
	"log"
	"server/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}
}

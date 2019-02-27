package main

import (
	"os"

	"github.com/MiguelMoll/joycast/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	web.Run(port)
}

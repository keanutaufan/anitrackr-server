package main

import (
	"github.com/keanutaufan/anitrackr-server/internal/app"
	"os"
)

func main() {
	server := app.NewServer()
	server.Logger.Fatal(server.Start(":" + os.Getenv("APP_PORT")))
}

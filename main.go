package main

import (
	"github.com/Ahmad940/health360/app"
	"github.com/Ahmad940/health360/platform/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connecting to database and initialize migrations
	db.InitializeMigration()

	// start server
	app.StartApp()
}

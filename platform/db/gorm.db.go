package db

import (
	"fmt"
	"os"

	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeMigration() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.GetEnv().DATABASE_URL), config.GormConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// migrations here
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Migration has failed: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app"
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

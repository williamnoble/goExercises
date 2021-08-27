package main

import (
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app"
	"github.com/williamnoble/goExercises/http/sql-gorm-api/api/app/config"
)

func main() {
	cfg := config.GetConfig()

	app := &app.Application{}
	app.Initialize(cfg)
	app.Run(":3000")
}

package main

import (
	"github.com/williamnoble/Projects/rest/SqlGormRestfulApi/api/app"
	"github.com/williamnoble/Projects/rest/SqlGormRestfulApi/api/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

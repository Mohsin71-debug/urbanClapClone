package main

import (
	"github.com/rsj-rishabh/urbanClapClone/server/app"
	"github.com/rsj-rishabh/urbanClapClone/server/config"
)

func main() {

	cfg := config.GetConfig()

	a := &app.App{}

	a.Initialize(cfg)

	a.InitializeDB()

	a.DBMigrate()

	a.Run(":8080")
}

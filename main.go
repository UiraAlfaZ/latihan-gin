package main

import (
	"latihan-gin/src/config"
	"latihan-gin/src/helper"
	"latihan-gin/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Router()
}

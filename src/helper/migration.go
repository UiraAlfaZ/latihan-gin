package helper

import (
	"latihan-gin/src/config"
	"latihan-gin/src/model"
)

func Migration() {
	config.DB.AutoMigrate(&model.Month{})
	config.DB.AutoMigrate(&model.Day{})
	config.DB.AutoMigrate(&model.Year{})
}

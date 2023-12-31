package database

import (
	"github.com/Beni-Zhang/task-5-pbi-btpns-Beny/app"
)

func Migrate() {
	DB.AutoMigrate(&app.User{}, &app.Photo{})
}
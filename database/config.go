package database

import (
   "gorm.io/driver/postgres"
   "gorm.io/gorm"
   "os"
)

var DB *gorm.DB

func InitDatabase() {
   dsn := os.Getenv("DB_HOST")

   database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   if err != nil {
      panic("Failed to connect to the database!")
   }

   DB = database
}
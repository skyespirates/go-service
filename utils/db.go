package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null" binding:"required"`
	Description string
	Completed   bool `gorm:"default:false"`
}

type User struct {
	gorm.Model
	Name string
}

var DB *gorm.DB

func InitDb() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:secret@tcp(127.0.0.1:3306)/go_service"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Task{}, &User{})
	Logger.Infow("Database connected")

}

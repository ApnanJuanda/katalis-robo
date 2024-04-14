package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"katalisRobo/component-store/helper"
	"os"
)

func NewDB() *gorm.DB {
	err := godotenv.Load("config/.env")
	helper.PanicIfError(err)

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYUSER"),
		os.Getenv("MYPASSWORD"),
		os.Getenv("MYHOST"),
		os.Getenv("MYPORT"),
		os.Getenv("MYDATABASE"))

	dialect := mysql.Open(mysqlInfo)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	fmt.Println("Success connect to DB")
	return db
}

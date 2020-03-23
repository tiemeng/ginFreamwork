package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	DB_NAME = "todo"
	DB_USER = "root"
	DB_PORT = "3306"
	DB_HOST = "localhost"
	DB_PASSWORD = "123456"
	)
var DBHelper *gorm.DB
var err error

func init(){
	dns := DB_USER+":"+DB_PASSWORD+"@("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local"
	DBHelper,err = gorm.Open("mysql",dns)
	if err!=nil{
		panic(err)
	}
	DBHelper.LogMode(true)
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)

}

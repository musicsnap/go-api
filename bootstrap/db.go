package bootstrap

import (
	"fmt"
	"go-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var WebDb *gorm.DB

func InitDB() {
	conf := config.GetConfig()

	host := conf.GetString("mysql.web.host")
	username := conf.GetString("mysql.web.username")
	password := conf.GetString("mysql.web.password")
	port := conf.GetString("mysql.web.port")
	dbname := conf.GetString("mysql.web.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	WebDb, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	dbPoolEditor, err := WebDb.DB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dbPoolEditor.SetMaxIdleConns(10)
	dbPoolEditor.SetMaxOpenConns(100)
	dbPoolEditor.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return WebDb
}

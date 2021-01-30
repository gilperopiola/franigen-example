package database

import (
	"fmt"
	"franigen-example/config"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDatabase() *gorm.DB {
	var err error
	dbname := config.Get().DBName
	dbhost := config.Get().DBHost
	dbport := config.Get().DBPort
	dbuser := config.Get().DBUsername
	dbpass := config.Get().DBPassword
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)

	args := url.Values{}
	args.Add("charset", "utf8mb4")
	args.Add("parseTime", "True")
	args.Add("loc", "Local")
	args.Add("timeout", "30s")
	argsString := args.Encode()

	db, err := gorm.Open("mysql", connectionString+"?"+argsString)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(20)
	db.DB().SetConnMaxLifetime(time.Hour * 5)

	return db
}

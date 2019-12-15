package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/BharathKumarRavichandran/cloud99/utils"
)

var DB *gorm.DB

func Open(config *utils.Config) error {

	var err error

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	// Connect to database
	conn, err := gorm.Open("mysql", connStr)
	if err != nil {
		utils.Logger.Fatal(err.Error())
		return err
	}

	// Output to stdout only if Debug is true
	if config.DEBUG {
		conn.LogMode(true)
	}

	// Globally setting DB as database connection
	DB = conn

	utils.Logger.Infof("Database connection successful on : %s", string(config.DB_PORT))
	return err
}

func Close() error {
	utils.Logger.Error("Database connection closed")
	return DB.Close()
}

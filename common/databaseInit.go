package common

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Masterminds/glide/path"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/config"
	"time"
)

var (
	DB_NAME,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT string
)

// Initialize the database configuration
func init() {
	//　loads the GO PATH of  current system　
	goPath := path.Gopath()
	fmt.Println(goPath)
	var buffer bytes.Buffer
	buffer.WriteString(goPath)
	buffer.WriteString("/src/Browser-achain/conf/databaseConfig.ini")
	c, _ := config.ReadDefault(buffer.String())
	DB_NAME, _ = c.String("DB", "DB_NAME")
	DB_USER, _ = c.String("DB", "DB_USER")
	DB_PASSWORD, _ = c.String("DB", "DB_PASSWORD")
	DB_HOST, _ = c.String("DB", "DB_HOST")
	DB_PORT, _ = c.String("DB", "DB_PORT")
	fmt.Println("\n The current database IP and port are:", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
}

func GetDbConnection() (db *sql.DB, err error) {

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))

	if err != nil {
		fmt.Println("\nconnection mysql error when open")
		return db, err
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	err = db.Ping()

	if err != nil {
		fmt.Println("\n init mysql connection when ping", err)
		return db, err
	}

	return db, nil
}

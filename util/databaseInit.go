package util

import (
	"github.com/robfig/config"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"time"
	"github.com/Masterminds/glide/path"
	"bytes"
)

var(
	DB_NAME,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT string
)

// 初始化数据库配置
func init()  {
//　路径是相对于当前文件的路径
	//c, _ := config.ReadDefault("../../config/databaseConfig.ini")
	gopath := path.Gopath()
	fmt.Println(gopath)
	var buffer bytes.Buffer
	buffer.WriteString(gopath)
	buffer.WriteString("/src/Browser-achain/conf/databaseConfig.ini")
	c, _ := config.ReadDefault(buffer.String())
	DB_NAME, _ = c.String("DB", "DB_NAME")
	DB_USER, _ = c.String("DB", "DB_USER")
	DB_PASSWORD, _ = c.String("DB", "DB_PASSWORD")
	DB_HOST, _ = c.String("DB", "DB_HOST")
	DB_PORT, _ = c.String("DB", "DB_PORT")
	fmt.Println("\n当前数据库ip以及端口为:",DB_HOST,DB_PORT,DB_USER,DB_PASSWORD,DB_NAME)
}

func GetDbConnection() (db *sql.DB,err error)  {


	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))


	if err != nil{
		fmt.Println("\nconnection mysql error when open")
		return db,err
	}
	db.SetConnMaxLifetime(time.Minute*5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	err = db.Ping()

	if err != nil{
		fmt.Println("\n init mysql connection when ping",err)
		return db,err
	}

	return db,nil
}


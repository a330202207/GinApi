package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	gorm.Model
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

//关闭数据库
func CloseDB() {
	defer db.Close()
}

//获取数据库
func GetDb() *gorm.DB {
	return db
}

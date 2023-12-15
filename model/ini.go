package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(str string) {
	db, err := gorm.Open("mysql", str)
	if err != nil {
		fmt.Println(err)
		panic("mysql数据库链接错误")
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//表名不加s
	db.SingularTable(true)
	//设置连接池
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	fmt.Println("数据库mysql链接成功")
	migration()
}

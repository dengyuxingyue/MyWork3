package main

import (
	"fmt"
	"strings"
	"work/cache"
	"work/conf"
	"work/model"

	"gopkg.in/ini.v1"
)

func Init() {
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	conf.ReadFromConf(file)
	//链接mysql
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	path := strings.Join([]string{conf.DbUser, ":", conf.DbPassWord, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	fmt.Printf("路径：%s\n", path)
	model.Database(path)
	cache.InitRedis(conf.RedisDbName, conf.RedisAddr, conf.RedisPw)

}

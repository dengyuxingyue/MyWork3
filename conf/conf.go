package conf

import (
	"gopkg.in/ini.v1"
)

var (
	HttpPort    string
	RedisAddr   string
	RedisDbName string
	RedisPw     string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
)

func ReadFromConf(file *ini.File) {
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}

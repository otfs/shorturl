package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Settings = new(Config)

type Config struct {
	Salt             string
	BaseUrl          string
	DbDriverName     string
	DbDataSourceName string
}

func Init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 短链
	Settings.BaseUrl = viper.GetString("shorturl.baseUrl")
	Settings.Salt = viper.GetString("shorturl.salt")

	// 数据库
	dbHost := viper.GetString(`db.host`)
	dbPort := viper.GetString(`db.port`)
	dbUser := viper.GetString(`db.user`)
	dbPass := viper.GetString(`db.password`)
	dbName := viper.GetString(`db.database`)
	Settings.DbDriverName = "mysql"
	Settings.DbDataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

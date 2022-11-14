package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config2 *Config2Struct

type MongoStruct struct {
	Host     string
	Port     string
	Username string
	Passwd   string
	Database string
}

type Config2Struct struct {
	Mongo         MysqlStruct
	MongoTemplate string
}

func GetConf2() *Config2Struct {
	return Config2
}

func ParseConf2() {
	v := viper.New()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	environment := os.Getenv("environment")
	absolutePath := filepath.Join(pwd, "conf2", environment+".yaml")
	v.SetConfigFile(absolutePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var config *Config2Struct
	err = v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	Config2 = config
}
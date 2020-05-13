package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName  string   `json:"app_name"`
	Port     string   `json:"port"`
	Mode     string   `json:"mode"`
	DataBase DataBase `json:"data_base"`
	Redis    Redis    `json:"redis"`
}

/*
* mysql 配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	DataBase string `json:"database"`
}

/*
* Redis 配置
 */
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"net_work"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}

var ServConfig AppConfig

// 初始化服务器配置
func InitConfig() *AppConfig {
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}

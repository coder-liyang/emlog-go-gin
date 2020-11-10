package config

import (
	"encoding/json"
	"os"
)

// Config 所有配置
type Config struct {
	Mysql `json:"mysql"`
}

//Mysql MySQL配置
type Mysql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Dbname   string `json:"dbname"`
	Charset  string `json:"charset"`
	Debug    bool   `json:"debug"`
	Prefix   string `json:"prefix"`
}

//获取所有配置
func GetConfigs() Config {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

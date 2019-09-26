package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
//所有配置
type Config struct {
	Mysql `json:"mysql"`
}
//MySQL配置
type Mysql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Dbname   string `json:"dbname"`
	Charset  string `json:"charset"`
	Debug    bool   `json:"debug"`
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

func GetConfig2() {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	config := Config{}
	conf := string(fd)
	err = json.Unmarshal([]byte(conf), &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Mysql.Dbname)
}

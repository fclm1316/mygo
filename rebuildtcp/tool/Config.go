package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	AppInfo Appcfg `json:"appinfo"`
	DbInfo  Dbcfg  `json:"dbinfo"`
}

type Appcfg struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

type Dbcfg struct {
	DbDriver string `json:"db_driver"`
	DbHost   string `json:"db_host"`
	DbUser   string `json:"db_user"`
	DbPwd    string `json:"db_password"`
	DbPort   string `json:"db_port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	Showsql  bool   `json:"showsql"`
}

var cnf *Config

func GetConfig() *Config {
	return cnf
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&cnf); err != nil {
		log.Panicln(err)
	}

	return cnf, nil
}

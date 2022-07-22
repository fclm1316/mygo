package tool

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DbInfo  Dbcnf
	AppInfo Appcfg
}

type Dbcnf struct {
	Dbip     string `yaml:"dbip"`
	Dbname   string `yaml:"dbname"`
	Dbuser   string `yaml:"dbuser"`
	Dbpasswd string `yaml:"dbpasswd"`
	Dbport   string `yaml:"dbport"`
	Charset  string `yaml:"charset"`
	Showsql  bool   `yaml:"showsql"`
}

type Appcfg struct {
	AppName string `yaml:"app_name"`
	AppMode string `yaml:"app_mode"`
	AppHost string `yaml:"app_host"`
	AppPort string `yaml:"app_port"`
}

var cnf *Config

func GetConfig() *Config {
	return cnf
}
func ParseConfig() (*Config, error) {
	filename := GetAllinoneConfPath()
	fileread, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("读取配置文件失败：%s", err)
	}

	if err = yaml.Unmarshal(fileread, &cnf); err != nil {
		log.Println("解析yaml失败：", err)
		return nil, err
	}

	return cnf, nil

}

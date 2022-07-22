package tool

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

var AllinoneDir string
var AllinoneLogPath string
var AllinoneConf string

const AllinoneLog = "allinone.log"
const AllinoneYaml = "app.yaml"

func GetCurrentDirectory() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Cannot get the process path", err)
	}
	AllinoneDir = dir
	return dir
}

func GetAllinoneLogFilePath() string {

	AllinoneLogPath = path.Join(GetCurrentDirectory(), "log", AllinoneLog)
	return AllinoneLog
}
func GetAllinoneConfPath() string {

	AllinoneConf = path.Join(GetCurrentDirectory(), "config", AllinoneYaml)
	return AllinoneConf
}

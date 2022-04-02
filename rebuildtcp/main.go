package main

import (
	"log"
	"mygo/rebuildtcp/controller"
	"mygo/rebuildtcp/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./conf/app.json")
	if err != nil {
		log.Panicln(err)
	}

	_, err = tool.OrmEngine(cfg)

	if err != nil {
		log.Panicln(err)
		return
	}

	app := gin.Default()

	registerRouter(app)
	app.Run(cfg.AppInfo.AppHost + ":" + cfg.AppInfo.AppPort)
}

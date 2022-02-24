package main

import "mygo/my/mylog"

func main() {
	log := mylog.NewLogger()
	log.Debug("aaaa")
	log.Info("bbbb")

}

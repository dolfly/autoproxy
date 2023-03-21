package cron

import (
	"runtime"

	"github.com/dolfly/autoproxy/internal/app"
	"github.com/jasonlvhit/gocron"
)

func Cron() {
	_ = gocron.Every(15).Minutes().Do(crawlTask)
	<-gocron.Start()
}

func crawlTask() {
	_ = app.InitConfigAndGetters("")
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}

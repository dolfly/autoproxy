package app

import (
	"github.com/jasonlvhit/gocron"
)

func Cron() {
	_ = gocron.Every(10).Minutes().Do(CrawlTGChannel)
	<-gocron.Start()
}

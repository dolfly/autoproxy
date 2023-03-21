package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dolfly/autoproxy/api"
	"github.com/dolfly/autoproxy/internal/app"
	"github.com/dolfly/autoproxy/internal/cron"
	"github.com/dolfly/autoproxy/internal/database"
	"github.com/dolfly/autoproxy/pkg/proxy"
)

var configFilePath = ""

func main() {
	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()
	if configFilePath == "" {
		configFilePath = os.Getenv("CONFIG_FILE")
	}
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := app.InitConfigAndGetters(configFilePath)
	if err != nil {
		panic(err)
	}
	database.InitTables()
	proxy.InitGeoIpDB()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	go cron.Cron()
	api.Run()
}

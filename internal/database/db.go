package database

import (
	"fmt"

	"github.com/dolfly/autoproxy/internal/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (err error) {
	fmt.Println(config.Config.Database.Driver)
	switch config.Config.Database.Driver {
	case "sqlite":
		dsn := "autoproxy.db"
		if url := config.Config.Database.Url; url != "" {
			dsn = url
		}
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			fmt.Println("DB connect success: ", DB.Name())
		}
	case "progres":
		// dsn := "user=proxypool password=proxypool dbname=proxypool port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		// if url := config.Config.Database.Url; url != "" {
		// 	dsn = url
		// }
		// if url := os.Getenv("DATABASE_PROGRES_URL"); url != "" {
		// 	dsn = url
		// }
		// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(logger.Silent),
		// })
		// if err == nil {
		// 	fmt.Println("DB connect success: ", DB.Name())
		// }
	case "mysql":
	default:

	}

	return
}

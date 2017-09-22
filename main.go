package main

import (
	"github.com/dzen-it/di/app"
	"github.com/dzen-it/di/common/configs"
	"github.com/dzen-it/di/system/db"
	"github.com/dzen-it/di/system/db/redis"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func main() {
	switch configs.Config.Mode {
	case "cache":
		db.Client = redis.NewClient(configs.Config.Redis, "", 0)
	case "redis":
		db.Client = redis.NewClient(configs.Config.Redis, "", configs.Config.RedisDB)
	default:
		log.Fatal("Undefined mode: ", configs.Config.Mode)
	}
	initLogger(log.DebugLevel)

	log.Info("Running di...")
	app.Start(configs.Config.Address)
}

func initLogger(level log.Level) {
	log.SetLevel(level)
	filenameHook := filename.NewHook(log.AllLevels...)
	log.AddHook(filenameHook)
}

package app

import (
	"cricradio-go-svc/Jobs"
	"cricradio-go-svc/db/kafka"
	"cricradio-go-svc/logger"
	"time"
)
import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Starting Cron Jobs")
	defer kafka.ControllerConn.Close()

	go LiveCron()
	go CommCron()

	logger.Info("About to start the application...")
	router.Run(":9900")
}

func LiveCron() {
	for {
		Jobs.LiveScraper()
		time.Sleep(1 * time.Minute)
	}
}

func CommCron() {
	for {
		Jobs.CommentaryScraper()
		time.Sleep(20 * time.Second)
	}
}

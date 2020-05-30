package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	log "github.com/sirupsen/logrus"
	fileController "gitlab.com/DockerDeployTest/src/controller"
	"os"
)

func registerRoutes() {
	r := gin.Default()
	r.Use(cors.Default())
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.SetOutput(os.Stdout)
	routeUploader(r)
	_ = r.Run()
}

func routeUploader(router *gin.Engine) {
	controller := fileController.NewFileUploader()
	router.POST("/ping", controller.UploadSingleFile)
	router.GET("/health", controller.Health)
}

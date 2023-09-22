package app

import (
	"github.com/EliasSantiago/uploader-s3/internal/handlers"
	conf "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Initialize(awsConfig *conf.AWSConfig) {
	gin.SetMode(gin.ReleaseMode)
	log := logger.GetLogger()

	r := gin.Default()

	handlers.SetupRoutes(r, awsConfig)

	err := r.Run(":8080")
	if err != nil {
		log.Error("Error starting server",
			zap.Error(err),
			zap.String("journey", "Initialize"),
		)
		panic(err)
	}
}

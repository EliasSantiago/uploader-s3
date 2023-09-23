package app

import (
	"strconv"

	"github.com/EliasSantiago/uploader-s3/internal/handlers"
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Initialize(cfg *config.Config, awsSession *s3.S3, bucketName string) {
	//gin.SetMode(gin.ReleaseMode)
	portNumber, err := strconv.Atoi(cfg.Port)
	if err != nil {
		logger.Error("Error parsing PORT environment variable", err,
			zap.String("journey", "Initialize"),
		)
		panic(err)
	}

	r := gin.Default()

	handlers.SetupRoutes(r, awsSession, bucketName)

	err = r.Run(":" + strconv.Itoa(portNumber))
	if err != nil {
		logger.Error("Error starting server on port", err,
			zap.String("journey", "Initialize"),
		)
		panic(err)
	}
}

package app

import (
	"github.com/EliasSantiago/uploader-s3/internal/handlers"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Initialize(awsSession *s3.S3, bucketName string) {
	gin.SetMode(gin.ReleaseMode)
	log := logger.GetLogger()

	r := gin.Default()

	handlers.SetupRoutes(r, awsSession, bucketName)

	err := r.Run(":8080")
	if err != nil {
		log.Error("Error starting server",
			zap.Error(err),
			zap.String("journey", "Initialize"),
		)
		panic(err)
	}
}

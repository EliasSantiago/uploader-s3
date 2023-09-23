package main

import (
	"github.com/EliasSantiago/uploader-s3/internal/app"
	"github.com/EliasSantiago/uploader-s3/internal/s3"
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"go.uber.org/zap"
)

func init() {
	err := logger.SetupLogger()
	if err != nil {
		panic("Error setting up logger: " + err.Error())
	}
}

func main() {
	awsConfig, err := config.LoadAWSConfig(".")
	if err != nil {
		panic(err)
	}

	awsSession, err := s3.InitAWS(awsConfig)
	if err != nil {
		logger.Error("Error initializing AWS session", err,
			zap.String("journey", "InitAWS"),
		)
		panic(err)
	}

	app.Initialize(awsSession, awsConfig.S3Bucket)
}

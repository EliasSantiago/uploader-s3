package main

import (
	"github.com/EliasSantiago/uploader-s3/internal/app"
	"github.com/EliasSantiago/uploader-s3/internal/s3"
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Error("Error loading config", err,
			zap.String("journey", "LoadConfig"),
		)
		panic(err)
	}

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

	app.Initialize(cfg, awsSession, awsConfig.S3Bucket)
}

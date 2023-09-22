package main

import (
	"github.com/EliasSantiago/uploader-s3/internal/app"
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
)

func init() {
	err := logger.SetupLogger()
	if err != nil {
		panic("Erro ao configurar o logger: " + err.Error())
	}
}

func main() {
	awsConfig, err := config.LoadAWSConfig(".")
	if err != nil {
		panic(err)
	}
	app.Initialize(awsConfig)
}

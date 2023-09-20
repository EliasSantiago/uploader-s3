package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/EliasSantiago/uploader-s3/configuration/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	AWS_REGION     = "AWS.REGION"
	AWS_ACCESS_KEY = "AWS.ACCESS_KEY"
	AWS_SECRET_KEY = "AWS.SECRET_KEY"
	AWS_S3_BUCKET  = "AWS.S3_BUCKET"

	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Read configs", err,
			zap.String("journey", "readInConfig"),
		)
		panic(err)
	}

	awsRegion := viper.GetString(AWS_REGION)
	accessKey := viper.GetString(AWS_ACCESS_KEY)
	secretKey := viper.GetString(AWS_SECRET_KEY)
	s3Bucket = viper.GetString(AWS_S3_BUCKET)

	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		})
	if err != nil {
		logger.Error("Init new session", err,
			zap.String("journey", "session.NewSession"),
		)
	}
	s3Client = s3.New(sess)
}

func main() {
	logger.Info("Starting uploader-s3",
		zap.String("journey", "startUploader"),
	)

	dir, err := os.Open("./tmp")
	if err != nil {
		logger.Error("Opening directory failed", err,
			zap.String("journey", "openDirectory"),
		)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Error("Error reading directory", err,
				zap.String("journey", "dir.Readdir"),
			)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFilename := fmt.Sprintf("./tmp/%s", filename)
	f, err := os.Open(completeFilename)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFilename)
		logger.Error("Error opening file", err,
			zap.String("journey", "os.Open"),
		)
		<-uploadControl
		errorFileUpload <- completeFilename
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		//fmt.Printf("Error uploading file %s: %s\n", completeFilename, err)
		logger.Error("Error uploading file", err,
			zap.String("journey", "s3Client.PutObject"),
		)
		<-uploadControl
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFilename)
	<-uploadControl
}

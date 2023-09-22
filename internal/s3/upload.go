package s3

import (
	"bytes"
	"fmt"

	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(filename string, uploadControl <-chan struct{}, c *gin.Context, fileBytes []byte, conf *config.AWSConfig, s3Client *s3.S3) error {
	defer func() {
		<-uploadControl
	}()

	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(conf.S3Bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		logger.Error("Error uploading file", err,
			zap.String("journey", "s3Client.PutObject"),
		)
		return err
	}
	fmt.Printf("File %s uploaded successfully\n", filename)
	return nil
}

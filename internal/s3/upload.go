package s3

import (
	"bytes"

	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(filename string, uploadControl <-chan struct{}, c *gin.Context, fileBytes []byte, awsSession *s3.S3, bucketName string) error {
	defer func() {
		<-uploadControl
	}()

	_, err := awsSession.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		logger.Error("Error uploading file",
			err,
			zap.String("journey", "s3Client.PutObject"),
		)
		return err
	}
	logger.Info("File uploaded successfully",
		zap.String("filename", filename),
		zap.String("bucket", bucketName),
	)
	return nil
}

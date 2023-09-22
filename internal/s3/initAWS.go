package s3

import (
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3Client *s3.S3

func InitAWS(awsConfig *config.AWSConfig) *s3.S3 {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(awsConfig.Region),
			Credentials: credentials.NewStaticCredentials(awsConfig.AccessKey, awsConfig.SecretKey, ""),
		})
	if err != nil {
		panic(err)
	}

	return s3.New(sess)
}

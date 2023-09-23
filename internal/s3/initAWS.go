package s3

import (
	config "github.com/EliasSantiago/uploader-s3/pkg/configs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	S3Client   *s3.S3
	awsSession *session.Session
)

func InitAWS(awsConfig *config.AWSConfig) (*s3.S3, error) {
	if S3Client != nil {
		return S3Client, nil
	}

	awsCfg := &aws.Config{
		Region:      aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKey, awsConfig.SecretKey, ""),
	}

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	S3Client = s3.New(sess)
	return S3Client, nil
}

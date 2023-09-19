package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	awsRegion := viper.GetString("aws.region")
	accessKey := viper.GetString("aws.access_key")
	secretKey := viper.GetString("aws.secret_key")

	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		})
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = viper.GetString("aws.s3_bucket")
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
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
			fmt.Printf("Error reading directory: %s\n", err)
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
	fmt.Printf("Uploadind file %s to bucket %s\n", completeFilename, s3Bucket)
	f, err := os.Open(completeFilename)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFilename)
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
		fmt.Printf("Error uploading file %s: %s\n", completeFilename, err)
		<-uploadControl
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFilename)
	<-uploadControl
}

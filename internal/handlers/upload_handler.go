package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"

	s3Upload "github.com/EliasSantiago/uploader-s3/internal/s3"
	"github.com/EliasSantiago/uploader-s3/pkg/logger"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	wg         sync.WaitGroup
	maxRetries = 3
)

type UploadRequest struct {
	Path  string                  `form:"path" binding:"required"`
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

func SetupRoutes(r *gin.Engine, awsSession *s3.S3, awsBucketName string) {
	r.POST("/api/v1/upload", func(c *gin.Context) {
		log := logger.GetLogger()

		uploadControl := make(chan struct{}, 100)
		var req UploadRequest

		if err := c.ShouldBind(&req); err != nil {
			log.Error("Error parsing request body",
				zap.Error(err),
				zap.String("journey", "ShouldBindJSON"),
			)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error parsing request body",
			})
			return
		}

		files := req.Files

		for _, file := range files {
			wg.Add(1)
			uploadControl <- struct{}{}

			go func(file *multipart.FileHeader) {
				defer wg.Done()
				uuidFilename := uuid.New().String()
				s3Path := normalizePath(req.Path) + uuidFilename

				for i := 0; i < maxRetries; i++ {
					fileBytes, err := s3Upload.RetryUpload(file)
					if err != nil {
						log.Error("Error reading file... %s\n",
							zap.Error(err),
							zap.String("journey", "RetryUpload"),
						)
						time.Sleep(2 * time.Second)
						continue
					}

					if err := s3Upload.UploadFile(s3Path, uploadControl, c, fileBytes, awsSession, awsBucketName); err == nil {
						return
					}
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Error uploading file %s", s3Path),
				})
			}(file)
		}
		wg.Wait()

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded successfully", len(files)),
		})
	})
}

func normalizePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}

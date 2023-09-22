package s3

import (
	"io"
	"mime/multipart"
)

func RetryUpload(file *multipart.FileHeader) ([]byte, error) {
	openFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	fileBytes, err := io.ReadAll(openFile)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

package service

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/matoous/go-nanoid"
	"log"
	"mime/multipart"
)

// this should be in a property file
var BUCKET_NAME = "gigamog.simple.server.2018"

func GetFileS3(record FileRecord, downloader *s3manager.Downloader) ([]byte, error) {
	buff := &aws.WriteAtBuffer{}

	numBytes, err := downloader.Download(buff,
		&s3.GetObjectInput{
			Bucket: aws.String(BUCKET_NAME),
			Key:    aws.String(record.FileId),
		})
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Println("downloaded file: ", numBytes)
	}

	return buff.Bytes(), nil
}

func SaveFileS3(bytes *bytes.Buffer, fileHeader multipart.FileHeader, uploader *s3manager.Uploader, record *FileRecord) error {
	id, _ := gonanoid.Nanoid()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(id),
		Body:   bytes,
	})

	if err != nil {
		log.Println("failed to upload file, %v", err)
		return err
	} else {

		record.setFileName(fileHeader.Filename)
		record.setFileSize(fileHeader.Size)
		record.setFileId(id)
		record.setFileLocation(result.Location)
	}
	return nil
}

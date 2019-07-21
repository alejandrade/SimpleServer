package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateS3BucketClient() (*s3manager.Uploader, *s3manager.Downloader) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	))

	return s3manager.NewUploader(sess), s3manager.NewDownloader(sess)
}

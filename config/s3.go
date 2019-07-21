package config

import (
	"SimpleServer/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateS3BucketClient(properties *util.Properties) (*s3manager.Uploader, *s3manager.Downloader) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(properties.Region)},
	))

	return s3manager.NewUploader(sess), s3manager.NewDownloader(sess)
}

package endpoints

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AppContext struct {
	DB           *dynamodb.DynamoDB
	S3Uploader   *s3manager.Uploader
	S3Downloader *s3manager.Downloader
}

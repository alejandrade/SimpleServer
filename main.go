package main

import (
	"SimpleServer/config"
)

func main() {
	amazonClient := config.CreateFileUploadIfNotExist()
	uploader, downloader := config.CreateS3BucketClient()
	config.HandleRequest(amazonClient, uploader, downloader)
}

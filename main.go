package main

import (
	"SimpleServer/config"
	"SimpleServer/util"
)

func main() {
	properties := util.LoadProperties()
	dbclient := config.CreateFileUploadIfNotExist(properties)
	uploader, downloader := config.CreateS3BucketClient(properties)
	config.HandleRequest(dbclient, uploader, downloader, properties)
}

package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type Properties struct {
	Database struct {
		TableName     string `json:"tableName"`
		ReadCapacity  int64  `json:"readCapacity"`
		WriteCapacity int64  `json:"writeCapacity"`
		Hash          struct {
			Key      string `json:"key"`
			DataType string `json:"dataType"`
		} `json:"hash"`
		Range struct {
			Key      string `json:"key"`
			DataType string `json:"dataType"`
		} `json:"range"`
	} `json:"database"`
	S3Bucket          string `json:"s3Bucket"`
	Region            string `json:"region"`
	UseCredentialFile bool   `json:"useCredentialFile"`
	AwsAccessKey      string `json:"awsAccessKey"`
	AwsSecretKey      string `json:"awsSecretKey"`
	Port              string `json:"port"`
}

func LoadProperties() *Properties {
	var config Properties
	configFile, err := os.Open("properties.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	setEnvs(config)
	return &config
}

func setEnvs(properties Properties) {
	if !properties.UseCredentialFile {
		os.Setenv("AWS_ACCESS_KEY_ID", properties.AwsAccessKey)
		os.Setenv("AWS_SECRET_ACCESS_KEY", properties.AwsSecretKey)
	}
}

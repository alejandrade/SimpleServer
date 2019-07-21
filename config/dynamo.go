package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

// all string values here should prob come out of a property file
func CreateFileUploadIfNotExist() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	//

	// Create DynamoDB client
	svc := dynamodb.New(sess, &aws.Config{
		Region: aws.String("us-east-2")},
	)

	tableName := "UserFileUpload"

	//I really should add an index here for User but I don't want to pay for AWS bill

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("FileId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("User"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("FileId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("User"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Println("Got error calling CreateTable:")
		log.Println(err.Error())
	} else {
		log.Println("Created the table", tableName)
	}

	return svc
}

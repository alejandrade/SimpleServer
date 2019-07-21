package config

import (
	"SimpleServer/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func CreateFileUploadIfNotExist(properties *util.Properties) *dynamodb.DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	//

	// Create DynamoDB client
	svc := dynamodb.New(sess, &aws.Config{
		Region: aws.String(properties.Region),
	})

	//I really should add an index here for User but I don't want to pay for AWS bill

	databaseProps := properties.Database

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(databaseProps.Hash.Key),
				AttributeType: aws.String(databaseProps.Hash.DataType),
			},
			{
				AttributeName: aws.String(databaseProps.Range.Key),
				AttributeType: aws.String(databaseProps.Range.DataType),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(databaseProps.Hash.Key),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String(databaseProps.Range.Key),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(databaseProps.ReadCapacity),
			WriteCapacityUnits: aws.Int64(databaseProps.WriteCapacity),
		},
		TableName: aws.String(databaseProps.TableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Println("Got error calling CreateTable:")
		log.Println(err.Error())
	} else {
		log.Println("Created the table", databaseProps.TableName)
	}

	return svc
}

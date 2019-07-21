package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"os"
)

var TABLE_NAME = "UserFileUpload"

func GetAllFilesDb(user string, db *dynamodb.DynamoDB) []FileRecord {
	filt := expression.Name("User").Equal(expression.Value(user))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TABLE_NAME),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	fileRecordItems := result.Items
	var fileRecords = make([]FileRecord, int(len(fileRecordItems)))

	for i, item := range fileRecordItems {
		fileRecord := FileRecord{}
		err = dynamodbattribute.UnmarshalMap(item, &fileRecord)
		fileRecords[i] = fileRecord

	}

	return fileRecords
}

func GetFileDb(fileId string, user string, db *dynamodb.DynamoDB) FileRecord {
	fileRecord := FileRecord{}
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"FileId": {
				S: aws.String(fileId),
			},
			"User": {
				S: aws.String(user),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return fileRecord
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &fileRecord)

	return fileRecord
}

func SaveFileDb(file FileRecord, db *dynamodb.DynamoDB) {
	av, err := dynamodbattribute.MarshalMap(file)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TABLE_NAME),
	}

	_, err = db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added '" + file.FileName)

}

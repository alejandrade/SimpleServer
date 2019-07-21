package service

import (
	"SimpleServer/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"log"
)

func GetAllFilesDb(user string, db *dynamodb.DynamoDB, properties *util.Properties) ([]FileRecord, error) {

	// normally here we would use query and add an index to dynamodb but I opted to use scan since this isn't a production project

	filt := expression.Name("User").Equal(expression.Value(user))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		log.Println("Got error building expression:", err.Error())
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(properties.Database.TableName),
	}

	result, err := db.Scan(params)
	if err != nil {
		log.Println("Query API call failed:")
		log.Println((err.Error()))
		return nil, err
	}

	fileRecordItems := result.Items
	var fileRecords = make([]FileRecord, int(len(fileRecordItems)))

	for i, item := range fileRecordItems {
		fileRecord := FileRecord{}
		err = dynamodbattribute.UnmarshalMap(item, &fileRecord)
		if err != nil {
			return nil, err
		}

		fileRecords[i] = fileRecord

	}

	return fileRecords, nil
}

func GetFileDb(fileId string, user string, db *dynamodb.DynamoDB, properties *util.Properties) (FileRecord, error) {
	fileRecord := FileRecord{}
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(properties.Database.TableName),
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
		log.Println(err.Error())
		return fileRecord, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &fileRecord)

	if err != nil {
		return fileRecord, err
	}

	return fileRecord, nil
}

func SaveFileDb(file FileRecord, db *dynamodb.DynamoDB, properties *util.Properties) error {
	av, err := dynamodbattribute.MarshalMap(file)
	if err != nil {
		log.Println("Got error marshalling new movie item:")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(properties.Database.TableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		log.Println("Got error calling PutItem:")
		return err
	}

	log.Println("Successfully added '", file.FileName)
	return nil

}

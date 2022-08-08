package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{Profile: "dynamodbclient"},
	))

	client := dynamodb.New(sess)

	putParam := &dynamodb.PutItemInput{
		TableName: aws.String("example-table"),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("1"),
			},
			"TIME": {
				S: aws.String(time.Now().String()),
			},
		},
	}

	_, err := client.PutItem(putParam)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	getParam := &dynamodb.GetItemInput{
		TableName: aws.String("example-table"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("1"),
			},
		},
	}

	res, err := client.GetItem(getParam)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Printf("current value is as follows\n %s", res)

}

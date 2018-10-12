package easydynamoclient

import (
	"fmt"
	"github.com/JuanValero25/EasyDynamoDB/dynamoclient"
	"github.com/JuanValero25/EasyDynamoDB/lambdaconfig"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

const ENVIRONMENT = "ENVIRONMENT"

type EasyDynamoClient struct {
	dynamoDbClient *dynamodb.DynamoDB
	queryHelper
}

func New() *EasyDynamoClient {

	easyClient := EasyDynamoClient{dynamoclient.NewDynamoClient(), queryHelper{}}

	return &easyClient
}

func NewWithTableObject(TableObject lambdaconfig.TableInfo) *EasyDynamoClient {

	easyClient := EasyDynamoClient{dynamoclient.NewDynamoClient(), queryHelper{TableObject}}

	return &easyClient
}

func (c EasyDynamoClient) Save(TableObject lambdaconfig.TableInfo) {

	ProcessTableInfo(TableObject)
	av, err := dynamodbattribute.MarshalMap(TableObject)

	if err != nil {
		fmt.Print("this is a shit")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(GetEnvironmentStage() + TableObject.TableName()),
	}

	output, err := c.dynamoDbClient.PutItem(input)
	fmt.Print(output)
	fmt.Print(err)
}

func (c EasyDynamoClient) Update(TableObject lambdaconfig.TableInfo) {

	input := updateReflectionHelper(TableObject)
	fmt.Println(input)
	response, err := c.dynamoDbClient.UpdateItem(input)
	fmt.Println(response)
	fmt.Println(err)

}

func (c *EasyDynamoClient) GetItemByHashKey(haskeyValue string, TableObject lambdaconfig.TableInfo) *lambdaconfig.TableInfo {
	getInput := c.queryByHasKey(haskeyValue)
	getInput.TableName = aws.String(GetEnvironmentStage() + TableObject.TableName())
	outputvalue, err := c.dynamoDbClient.GetItem(getInput)
	fmt.Println(getInput)
	fmt.Println(outputvalue)
	fmt.Println(err)

	err = dynamodbattribute.UnmarshalMap(outputvalue.Item, TableObject)

	if err != nil {
		fmt.Println("Failed to unmarshal Record, %v", err)
	}
	return &TableObject

}

func GetEnvironmentStage() string {
	return os.Getenv(ENVIRONMENT)
}

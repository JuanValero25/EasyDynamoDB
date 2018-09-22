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
}

func New() *EasyDynamoClient {

	easyClient := EasyDynamoClient{dynamoclient.NewDynamoClient()}

	return &easyClient
}

func (c EasyDynamoClient) Save(tableObject lambdaconfig.TableInfo) {

	ProcessTableInfo(tableObject)

	av, err := dynamodbattribute.MarshalMap(tableObject)

	if err != nil {
		fmt.Print("this is a shit")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(GetEnvironmentStage() + tableObject.TableName()),
	}

	output, err := c.dynamoDbClient.PutItem(input)
	fmt.Print(output)
	fmt.Print(err)
}

func GetEnvironmentStage() string {
	return os.Getenv(ENVIRONMENT)
}

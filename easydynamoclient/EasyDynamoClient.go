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
	response,err :=c.dynamoDbClient.UpdateItem(input)
	fmt.Println(response)
	fmt.Println(err)



}

func GetEnvironmentStage() string {
	return os.Getenv(ENVIRONMENT)
}

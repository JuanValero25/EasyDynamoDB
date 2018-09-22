package dynamoclient

import (
	"github.com/JuanValero25/EasyDynamoDB/lambdaconfig"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoClient() *dynamodb.DynamoDB {

	return dynamodb.New(lambdaconfig.GetConfig())
}

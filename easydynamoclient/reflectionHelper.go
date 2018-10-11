package easydynamoclient

import (
	"github.com/JuanValero25/EasyDynamoDB/lambdaconfig"
	"github.com/satori/go.uuid"
	"reflect"
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ProcessTableInfo(tableObject lambdaconfig.TableInfo) {

	t := reflect.TypeOf(tableObject).Elem()

	for i := 0; i < t.NumField(); i++ {
		_, trued := t.Field(i).Tag.Lookup("AutoGenerated")
		if trued {
			ui, _ := uuid.NewV4()
			reflect.ValueOf(tableObject).Elem().Field(i).SetString(ui.String())

		}
	}

}

func updateReflectionHelper(tableObject lambdaconfig.TableInfo) *dynamodb.UpdateItemInput {

	start := time.Now()
	av, _ := dynamodbattribute.MarshalMap(tableObject)
    delete(av,"AccountId")

	t := reflect.TypeOf(tableObject)
	v:= reflect.ValueOf(tableObject)
	startSet:="set "
	var keyValue map[string]*dynamodb.AttributeValue
	for i := 0; i < t.NumField(); i++ {
		if  v.Field(i).String() != ""{
			fmt.Println()
			startSet=startSet+t.Field(i).Name+" = :"+t.Field(i).Name+", "
		}

		_, trued := t.Field(i).Tag.Lookup("AutoGenerated")
		if trued {
			keyValue=map[string]*dynamodb.AttributeValue{t.Field(i).Name: {
				S: aws.String(v.Field(i).String()),
			}}

		}
	}

	startSet = startSet[:len(startSet)-2]
	startSet= "set :InterestingData = :i"
	input := &dynamodb.UpdateItemInput{
		Key:keyValue,
		ExpressionAttributeValues: av,
		UpdateExpression: &startSet,
		TableName:                 aws.String(GetEnvironmentStage() + tableObject.TableName()),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}
	fmt.Println(startSet)
	fmt.Println(keyValue)

	elapsed := time.Since(start)
	fmt.Println("reflections time took", elapsed)

	return input
}

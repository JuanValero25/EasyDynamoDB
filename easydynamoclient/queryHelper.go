package easydynamoclient

import (
	"github.com/JuanValero25/EasyDynamoDB/lambdaconfig"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
)

type queryHelper struct {
	TableObject lambdaconfig.TableInfo
}

const (
	dynamoAttribute = "dynamoAttribute"
	dynamoHasKey    = "dynamoHasKey"
	dynamoSortKey   = "dynamoSortKey"
	dynamoindextKey = "dynamoindextKey"
)

func (r queryHelper) queryByHasKey(hashKeyString string) *dynamodb.GetItemInput {
	getQueryInput := &dynamodb.GetItemInput{Key: map[string]*dynamodb.AttributeValue{
		r.getHaskeyName(): {S: aws.String(hashKeyString)},
	},
		TableName: aws.String(r.TableObject.TableName()),
	}
	return getQueryInput
}

func (r queryHelper) getHaskeyName() string {

	t := reflect.TypeOf(r.TableObject)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		haskeyName, isHaskey := field.Tag.Lookup(dynamoHasKey)
		if isHaskey {
			return haskeyName
		}
	}
	return ""
}

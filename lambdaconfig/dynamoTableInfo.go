package lambdaconfig

import "reflect"

type TableInfo interface {
	TableName() string
	ReflectedValue() reflect.Value
}

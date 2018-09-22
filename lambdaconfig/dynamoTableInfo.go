package lambdaconfig

type TableInfo interface {
	TableName() string
	CanAddr() bool
}

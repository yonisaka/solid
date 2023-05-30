package entity

// EntityInterface is an interface
type EntityInterface interface {
	TableName() string
}

// Entity is a struct
type Entity struct {
	Entity interface{}
}

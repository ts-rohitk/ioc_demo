package builder

import "go.mongodb.org/mongo-driver/bson"

type SetQueryBuilder struct {
	fields bson.M
}

func NewSetQueryBuilder() *SetQueryBuilder {
	return &SetQueryBuilder{
		fields: make(bson.M),
	}
}

func (sqb *SetQueryBuilder) AddForUpdate(fieldName string, value any) *SetQueryBuilder {
	sqb.fields[fieldName] = value
	return sqb
}

func (sqb *SetQueryBuilder) Build() bson.M {
	return sqb.fields
}

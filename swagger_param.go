package ginHelper

import (
	"reflect"

	"github.com/go-openapi/spec"
)

func queryParams(typeOf reflect.Type) []spec.Parameter {
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	params := []spec.Parameter{}
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		field := typeOf.FieldByIndex([]int{i})
		if field.Type.Kind() == reflect.Struct {
			params = append(params, queryParams(field.Type)...)
			continue
		}
		formName := field.Tag.Get("form")
		if formName != "" {
			params = append(params, *spec.QueryParam(formName))
		}
	}
	return params
}

func RjsonSchema(param interface{}) *spec.Schema {
	return &spec.Schema{}
}

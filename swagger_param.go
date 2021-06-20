package ginHelper

import (
	"fmt"
	"path"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-openapi/spec"
)

func pathParams(sp *SwaggerPath) (params []spec.Parameter) {
	paths := strings.Split(sp.Path, "/")
	fmt.Println(paths)
	newPath := "/"
	for _, p := range paths {
		if len(p) <= 1 || p[0] != ':' {
			newPath = path.Join(newPath, p)
			continue
		}
		// fmt.Println(len(p))

		params = append(params, *spec.PathParam(p[1:]))
		newPath = path.Join(newPath, "{"+p[1:]+"}")
		continue

	}
	sp.Path = newPath
	return params
}

func queryParams(typeOf reflect.Type) []spec.Parameter {
	typeOf = typeElem(typeOf)
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

func JsonSchemas(typeOf reflect.Type) (schema *spec.Schema) {

	typeOf = typeElem(typeOf)

	switch typeOf.Kind() {
	case reflect.Struct:
		return kindStruct2Schema(typeOf)
	case reflect.Bool:
		return spec.BoolProperty()
	case reflect.Int, reflect.Uint:
		return spec.Int64Property()
	case reflect.Int8, reflect.Uint8:
		return spec.Int8Property()
	case reflect.Int16, reflect.Uint16:
		return spec.Int16Property()
	case reflect.Int32, reflect.Uint32:
		return spec.Int32Property()
	case reflect.Int64, reflect.Uint64:
		return spec.Int64Property()
	case reflect.Float32:
		return spec.Float32Property()
	case reflect.Float64:
		return spec.Float64Property()
	case reflect.Slice, reflect.Array:
		return kindArray2Schema(typeOf)
	case reflect.String:
		return spec.StringProperty()
	default:
		fmt.Println("unknow", typeOf)
		return &spec.Schema{}
	}
}
func kindArray2Schema(typeOf reflect.Type) *spec.Schema {
	return spec.ArrayProperty(JsonSchemas(typeOf.Elem()))
}
func kindStruct2Schema(typeOf reflect.Type) *spec.Schema {
	schema := &spec.Schema{SchemaProps: spec.SchemaProps{Type: []string{"object"},
		AdditionalProperties: nil}}
	// fieldNum := typeOf.NumField()
	fields := getAllField(typeOf)
	for len(fields) > 0 {
		field := fields[0]
		fields = fields[1:]
		name := field.Tag.Get("json")
		if name == "" {
			name = lcfirst(field.Name)
		}
		if name == "-" {
			continue
		}
		if field.Anonymous {
			if field.Type.Kind() != reflect.Struct {
				fields = append(fields, reflect.StructField{
					Name:      field.Name,
					PkgPath:   field.PkgPath,
					Type:      field.Type,
					Tag:       field.Tag,
					Offset:    field.Offset,
					Index:     field.Index,
					Anonymous: false,
				})
			} else {
				fields = append(fields, getAllField(field.Type)...)
			}
			continue
		}
		schema.SetProperty(name, *JsonSchemas(field.Type))
	}
	return schema
}

func kindBool2Schema(typeOf reflect.Type) *spec.Schema {
	return spec.BoolProperty()
}

func typeElem(typeOf reflect.Type) reflect.Type {
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	return typeOf
}

func getAllField(typeOf reflect.Type) []reflect.StructField {
	typeOf = typeElem(typeOf)

	fields := []reflect.StructField{}

	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		field := typeOf.FieldByIndex([]int{i})
		fields = append(fields, field)
	}
	return fields
}

func lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

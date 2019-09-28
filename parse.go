package gostruct2graphql

import (
	"errors"
	"reflect"

	"github.com/graphql-go/graphql"
)

// ParseStructList transforms a list of structs into a graphql object list
func ParseStructList(list interface{}) (*graphql.List, error) {
	inType := reflect.TypeOf(list)

	if inType.Kind() != reflect.Slice || inType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("Input is not a slice of structs")
	}

	name, fields, err := innerParseStruct(inType.Elem(), map[string]uint8{})
	return graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})), err
}

// ParseStruct transforms a struct into a graphql object
func ParseStruct(obj interface{}) (*graphql.Object, error) {
	inType := reflect.TypeOf(obj)

	if inType.Kind() != reflect.Struct {
		return nil, errors.New("Input is not a struct")
	}

	name, fields, err := innerParseStruct(inType, map[string]uint8{})
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	}), err
}

// innerParseStruct parses a struct
// The usedFields is to track code loops
func innerParseStruct(t reflect.Type, usedFields map[string]uint8) (string, *graphql.Fields, error) {
	fields := graphql.Fields{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		name, fieldType, err := parseField(field.Type, usedFields, field.Name)
		if err != nil {
			return "", nil, err
		}

		fields[name] = &graphql.Field{
			Type: fieldType,
		}
	}

	return t.Name(), &fields, nil
}

func parseField(t reflect.Type, usedFields map[string]uint8, alterName ...string) (string, graphql.Type, error) {
	name := t.Name()
	if len(alterName) > 0 {
		name = alterName[0]
	}

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return name, graphql.Int, nil
	case reflect.Float32, reflect.Float64:
		return name, graphql.Float, nil
	case reflect.String:
		return name, graphql.String, nil
	case reflect.Struct:
		_, innerFields, err := innerParseStruct(t, usedFields)
		if err != nil {
			return "", nil, err
		}
		return name, graphql.NewObject(graphql.ObjectConfig{
			Name:   name,
			Fields: innerFields,
		}), nil
	case reflect.Slice:

		return "", nil, nil
	default:
		return "", nil, errors.New("Unkown field kind: " + t.Kind().String())
	}
}

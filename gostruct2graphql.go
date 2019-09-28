package gostruct2graphql

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
)

func describeStruct(s interface{}) (graphql.Fields, error) {
	fields := graphql.Fields{}

	if reflect.ValueOf(s).Kind() == reflect.Slice {
		switch reflect.ValueOf(s).Type().Elem().Kind() {
		case reflect.Struct:

			iType := reflect.TypeOf(s).Elem()
			return describeSlice(iType)

		default:
			return nil, errors.New("something wrong happend")
		}

	}

	iValue := reflect.ValueOf(s)
	iType := reflect.TypeOf(s)

	for i := 0; i < iType.NumField(); i++ {
		v := iValue.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			innerFields, err := describeStruct(v.Interface())
			if err != nil {
				return nil, err
			}

			fields[iType.Field(i).Name] = &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name:   iType.Field(i).Name,
					Fields: innerFields,
				}),
			}
		case reflect.Slice:
			innerFields, err := describeStruct(v.Interface())
			if err != nil {
				return nil, err
			}

			fields[iType.Field(i).Name] = &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name:   iType.Field(i).Name,
					Fields: innerFields,
				})),
			}

		default:
			describeSimpleType(iType.Field(i).Name, iType.Field(i).Type.String(), string(iType.Field(i).Tag))
		}
	}

	return fields, nil
}

func describeSlice(iType reflect.Type) (graphql.Fields, error) {
	fields := graphql.Fields{}

	for i := 0; i < iType.NumField(); i++ {
		switch iType.Field(i).Type.Kind() {
		case reflect.Struct:
			innerFields, err := describeSlice(iType.Field(i).Type)
			if err != nil {
				return nil, err
			}

			fields[iType.Field(i).Name] = &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name:   iType.Field(i).Name,
					Fields: &innerFields,
				}),
			}
		case reflect.Slice:
			innerFields, err := describeSlice(iType.Field(i).Type.Elem())
			if err != nil {
				return nil, err
			}

			fields[iType.Field(i).Name] = &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name:   iType.Field(i).Name,
					Fields: &innerFields,
				})),
			}
		default:
			field, name, err := describeSimpleType(iType.Field(i).Name, iType.Field(i).Type.String(), string(iType.Field(i).Tag))
			if err != nil {
				return nil, err
			}
			fields[name] = field
		}
	}

	fmt.Println(fields)

	return fields, nil
}

func describeSimpleType(text ...interface{}) (field *graphql.Field, fieldName string, err error) {
	fieldName = text[0].(string)

	switch text[1] {
	case "string":
		field = &graphql.Field{
			Type: graphql.String,
		}
	case "bool":
		field = &graphql.Field{
			Type: graphql.Boolean,
		}
	case "uint64":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "uint32":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "uint16":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "uint8":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "uint":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "int64":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "int32":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "int16":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "int8":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	case "int":
		field = &graphql.Field{
			Type: graphql.Int,
		}
	default:
		err = errors.New("no such format")
		return
	}

	return
}

// GetRootDescription returns the root description of a Struct
func GetRootDescription(strct interface{}) (*graphql.Object, error) {
	if reflect.TypeOf(strct).Kind() != reflect.Struct {
		return nil, errors.New("Input is not a struct")
	}

	fields, err := describeStruct(strct)

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   reflect.TypeOf(strct).Name(),
		Fields: &fields,
	}), err
}

// GetRootDescriptionSlice returns the root description of a Slice
func GetRootDescriptionSlice(slc interface{}) (*graphql.List, error) {
	if reflect.TypeOf(slc).Kind() != reflect.Slice {
		return nil, errors.New("Input is not a slice")
	}

	fields, err := describeStruct(slc)
	return graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
		Name:   reflect.TypeOf(slc).Name(),
		Fields: &fields,
	})), err
}

// Test is is
var Test = graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
	Name: "Repository",
	Fields: graphql.Fields{
		"CacheMaxSeconds": &graphql.Field{
			Type: graphql.Int,
		},
		"CurrentTime": &graphql.Field{
			Type: graphql.Int,
		},
		"Doc": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "Doc",
				Fields: graphql.Fields{
					"TropData": &graphql.Field{
						Type: graphql.NewObject(graphql.ObjectConfig{
							Name: "TropData",
							Fields: graphql.Fields{
								"Two016": &graphql.Field{
									Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
										Name: "Two016",
										Fields: graphql.Fields{
											"Active": &graphql.Field{
												Type: graphql.Boolean,
											},
											"Category": &graphql.Field{
												Type: graphql.String,
											},
											"Status": &graphql.Field{
												Type: graphql.String,
											},
											"TropID": &graphql.Field{
												Type: graphql.String,
											},
											"TropName": &graphql.Field{
												Type: graphql.String,
											},
										},
									})),
								},
							},
						}),
					},
					"TropHdr": &graphql.Field{
						Type: graphql.NewObject(graphql.ObjectConfig{
							Name: "TropHdr",
							Fields: graphql.Fields{
								"TNum": &graphql.Field{
									Type: graphql.Int,
								},
							},
						}),
					},
				},
			}),
		},
		"GeneratedTime": &graphql.Field{
			Type: graphql.Int,
		},
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"Status": &graphql.Field{
			Type: graphql.Int,
		},
	},
}))

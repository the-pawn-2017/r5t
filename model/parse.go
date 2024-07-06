package model

import (
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
)

// need to debug
func ParseModel(t reflect.Type, s *openapi3.Schema) {
	switch t.Kind() {
	case reflect.Bool:
		s.Type = &openapi3.Types{openapi3.TypeBoolean}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s.Type = &openapi3.Types{openapi3.TypeInteger}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.Type = &openapi3.Types{openapi3.TypeInteger}
	case reflect.Float32, reflect.Float64:
		s.Type = &openapi3.Types{openapi3.TypeNumber}
	case reflect.String:
		s.Type = &openapi3.Types{openapi3.TypeString}
	case reflect.Pointer:
		vInStruct := new(openapi3.Schema)
		ParseModel(t.Elem(), vInStruct)
		s.Type = vInStruct.Type
	case reflect.Struct:
		s.Type = &openapi3.Types{openapi3.TypeObject}
		// every item should deal
		for i := 0; i < t.NumField(); i++ {
			// jump out private
			if t.Field(i).PkgPath != "" {
				continue
			}
			tag, have := t.Field(i).Tag.Lookup("json")
			if s.Properties == nil {
				s.Properties = openapi3.Schemas{}
			}
			if have {
				s.Title = tag
			} else {
				s.Title = t.Field(i).Name
			}
			vInStruct := new(openapi3.Schema)
			ParseModel(t.Field(i).Type, vInStruct)
			s.Properties[s.Title] = &openapi3.SchemaRef{
				Value: vInStruct,
			}

		}
	// case reflect.Map 缓一下
	case reflect.Array, reflect.Slice:
		s.Type = &openapi3.Types{openapi3.TypeArray}
		s.Items = new(openapi3.SchemaRef)
		vInStruct := new(openapi3.Schema)
		ParseModel(t.Elem(), vInStruct)
		s.Items.Value = vInStruct
	}

}
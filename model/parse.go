package model

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

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
		if s.Properties == nil {
			s.Properties = openapi3.Schemas{}
		}
		// every item should deal
		for i := 0; i < t.NumField(); i++ {
			// jump out private
			if t.Field(i).PkgPath != "" {
				continue
			}
			// embed struct
			if t.Field(i).Anonymous {
				ParseModel(t.Field(i).Type, s)
				continue
			}
			tag, hasCustomName := t.Field(i).Tag.Lookup("json")
			var realName string
			if hasCustomName {
				realName = tag
			} else {
				realName = t.Field(i).Name
			}
			vInStruct := new(openapi3.Schema)
			// deal validate
			validate, hasValidate := t.Field(i).Tag.Lookup("validate")
			if hasValidate {
				IsRequired := parseValidate(validate, vInStruct)
				if IsRequired {
					s.Required = append(s.Required, realName)
				}
			}
			ParseModel(t.Field(i).Type, vInStruct)
			s.Properties[realName] = &openapi3.SchemaRef{
				Value: vInStruct,
			}

		}
	// case reflect.Map now, do not support map, need dev
	case reflect.Array, reflect.Slice:
		s.Type = &openapi3.Types{openapi3.TypeArray}
		s.Items = new(openapi3.SchemaRef)
		vInStruct := new(openapi3.Schema)
		ParseModel(t.Elem(), vInStruct)
		s.Items.Value = vInStruct
	}

}

// the "openAPI required validate" put the required config in parent's Schema,Not it self, So need return
func parseValidate(validateStr string, s *openapi3.Schema) bool {
	var requiredField = false
	for _, v := range strings.Split(validateStr, ",") {

		vSplited := strings.Split(v, "=")[:]
		if len(vSplited) > 2 {
			requiredField = false
			continue
		}
		switch vSplited[0] {
		case "required":
			requiredField = true
		case "gte":
			re, _ := strconv.ParseFloat(vSplited[1], 64)
			s.Min = &re
		case "lte":
			re, _ := strconv.ParseFloat(vSplited[1], 64)
			s.Max = &re
		case "oneof":
			for _, v := range strings.Split(vSplited[1], " ") {
				s.Enum = append(s.Enum, v)
			}
		}

	}
	return requiredField
}

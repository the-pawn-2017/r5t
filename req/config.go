package req

import (
	"github.com/the-pawn-2017/r5t/header"

	"github.com/getkin/kin-openapi/openapi3"
)

type ReqModelOpts func(s *openapi3.RequestBody)

// about request
func FormFile(name string, description string, required bool) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		for k, v := range s.Content {
			switch k {
			case header.MultipartFormData:
				if required {
					if v.Schema.Value.Required == nil {
						v.Schema.Value.Required = make(openapi3.Types, 0)
					}
					alreadyHas := false
					for _, v := range v.Schema.Value.Required {
						if v == name {
							alreadyHas = true
							break
						}

					}
					if !alreadyHas {
						v.Schema.Value.Required = append(v.Schema.Value.Required, name)
					}
				}
				if v.Schema.Value.Properties == nil {
					v.Schema.Value.Properties = make(openapi3.Schemas)
				}
				v.Schema.Value.Properties[name] = &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:        &openapi3.Types{"string"},
						Format:      "bin",
						Description: description,
					},
				}
				// maybe has png、GiF, so using switch
			}
		}
	}
}

func Required() ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Required = true
	}
}

func Desc(description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Description = description
	}
}

func Example[T any](example T) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		for _, v := range s.Content {
			if v != nil {
				v.Example = example
			}

		}
	}
}
func Default[T any](v T) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		for _, v := range s.Content {
			if v != nil {
				v.Schema.Value.Default = v
			}

		}
	}
}

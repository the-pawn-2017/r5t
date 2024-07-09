package req

import (
	"r5t/header"

	"github.com/getkin/kin-openapi/openapi3"
)

type ReqModelOpts func(s *openapi3.RequestBody)

// about request
func WithForm(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[header.MultipartFormData] = nil
	}
}

// Deprecated: use WithOther instead.
func WithJSON(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[header.ApplicationJson] = nil
	}
}

func WithRequired(required bool) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Required = true
	}
}

func WithDesc(description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Description = description
	}
}

func WithExample[T any](example T) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		for _, v := range s.Content {
			if v != nil {
				v.Example = example
			}

		}
	}
}

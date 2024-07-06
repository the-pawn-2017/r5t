package model

import (
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
)

// Model is a model used in one or more routes.
type ModelOpts func(s *openapi3.Schema)

// some function
func WithModelDesc(desc string) ModelOpts {
	return func(s *openapi3.Schema) {
		s.Description = desc
	}
}

type Model struct {
	Type    reflect.Type
	Options []ModelOpts
}

// ModelOf creates a model of type T.
func ModelOf[T any]() Model {
	var t T
	m := Model{
		Type: reflect.TypeOf(t),
	}
	return m
}

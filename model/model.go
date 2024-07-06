package model

import (
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
)

const ReqJSON = "application/json"
const ReqXML = "application/xml"

// need dev
const ReqForm = "application/xxxxxx"

// Model is a model used in one or more routes.
type ModelOpts func(s *openapi3.Schema)

// some function
func WithModelDesc(desc string) ModelOpts {
	return func(s *openapi3.Schema) {
		s.Description = desc
	}
}

type ReqModelOpts func(s *openapi3.RequestBody)

// about request
func WithReqForm(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[ReqForm] = nil
	}
}
func WithReqJSON(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[ReqJSON] = nil
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

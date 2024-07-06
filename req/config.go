package req

import "github.com/getkin/kin-openapi/openapi3"

type ReqModelOpts func(s *openapi3.RequestBody)

const ReqJSON = "application/json"
const ReqXML = "application/xml"

// need dev
const ReqForm = "application/xxxxxx"

// about request
func WithForm(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[ReqForm] = nil
	}
}
func WithJSON(required bool, description string) ReqModelOpts {
	return func(s *openapi3.RequestBody) {
		s.Content[ReqJSON] = nil
	}
}

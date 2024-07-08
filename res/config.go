package res

import "github.com/getkin/kin-openapi/openapi3"

type ResModelOpts func(s *openapi3.Response)

const ReqJSON = "application/json"
const ResXML = "application/xml"

// need dev
const ReqForm = "application/xxxxxx"

// about request
func WithForm(required bool, description string) ResModelOpts {
	return func(s *openapi3.Response) {
		s.Content[ReqJSON] = nil
	}
}

/* func WithJSON(required bool, description string) ResModelOpts {
	return func(s *openapi3.Response) {
		s.Content[ReqJSON] = nil
	}
} */

func WithHeader() ResModelOpts {
	return func(s *openapi3.Response) {
	}
}

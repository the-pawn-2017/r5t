package res

import (
	"github.com/the-pawn-2017/r5t/header"

	"github.com/getkin/kin-openapi/openapi3"
)

type ResModelOpts func(s *openapi3.Response)

const ReqJSON = "application/json"
const ResXML = "application/xml"

// need dev
const ReqForm = "application/xxxxxx"

// about request
func Form(required bool, description string) ResModelOpts {
	return func(s *openapi3.Response) {
		s.Content[ReqJSON] = nil
	}
}

func Example(e any) ResModelOpts {
	return func(s *openapi3.Response) {
		s.Content[header.ApplicationJson].Example = e
	}
}

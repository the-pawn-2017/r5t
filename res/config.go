package res

import (
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

		for k, v := range s.Content {
			if v != nil {
				s.Content[k].Example = e
			}
		}
	}
}

func Desc(e string) ResModelOpts {
	return func(s *openapi3.Response) {
		s.Description = &e
	}
}

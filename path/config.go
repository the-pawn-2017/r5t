package path

import "github.com/getkin/kin-openapi/openapi3"

type PathOpts func(s *openapi3.Operation)

func WithDesc(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc

	}
}

func WithSummary(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc
	}
}

func WithTags(tags []string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Tags = tags
	}
}

func WithSecurity() PathOpts {
	return func(s *openapi3.Operation) {
		// s.Description = desc
	}
}

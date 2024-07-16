package path

import "github.com/getkin/kin-openapi/openapi3"

type PathOpts func(s *openapi3.Operation)

func Desc(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc

	}
}

func Summary(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc
	}
}

func Tags(tags ...string) PathOpts {
	return func(s *openapi3.Operation) {
		if s.Tags == nil {
			s.Tags = make([]string, 0)
		}
		s.Tags = append(s.Tags, tags...)
	}
}

func Security(tokenName string, auth []string) PathOpts {
	return func(s *openapi3.Operation) {
		if s.Security == nil {
			s.Security = openapi3.NewSecurityRequirements()
		}
		*s.Security = append(*s.Security, openapi3.SecurityRequirement{
			tokenName: auth,
		})
	}
}

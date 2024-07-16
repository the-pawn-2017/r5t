package spec

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type SpecOpts func(s *openapi3.T)

// contact
func Contact(name string, mail string, url string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Contact = &openapi3.Contact{
			Name:  name,
			Email: mail,
			URL:   url,
		}
	}
}

// title
func Title(title string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Title = title
	}
}

// version
func Version(v string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Version = v
	}
}

// servers
func Server(url string, desc string) SpecOpts {
	return func(s *openapi3.T) {
		s.Servers = append(s.Servers, &openapi3.Server{
			URL:         url,
			Description: desc,
		})
	}
}

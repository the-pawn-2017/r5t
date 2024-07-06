package spec

import (
	"r5t/api"

	"github.com/getkin/kin-openapi/openapi3"
)

type spec struct {
	// some info, do not have apis
	root openapi3.T
}

func (s *spec) MarshalJSON() ([]byte, error) {
	return s.root.MarshalJSON()
}

type SpecOpts func(s *openapi3.Info)

// contact
func WithContact(name string, mail string, url string) SpecOpts {
	return func(s *openapi3.Info) {
		s.Contact = &openapi3.Contact{
			Name:  name,
			Email: mail,
			URL:   url,
		}
	}
}

// title
func WithTitle(title string) SpecOpts {
	return func(s *openapi3.Info) {
		s.Title = title
	}
}

// version
func WithVersion(v string) SpecOpts {
	return func(s *openapi3.Info) {
		s.Version = v
	}
}

// 新建一个Spec
func NewSpec(specs ...SpecOpts) *spec {
	var s spec
	s.root.Info = &openapi3.Info{}
	for _, v := range specs {
		v(s.root.Info)
	}
	s.root.OpenAPI = "3.0.0"
	return &s
}

func (s *spec) addNewApi(path string, method string, opts []api.PathOpts) *api.API {
	var newApi *api.API = &api.API{
		Operation: &openapi3.Operation{},
	}
	s.root.AddOperation(path, method, newApi.Operation)
	newApi.DealPathItem(newApi.Operation, opts)
	return newApi
}

// some options function
func (s *spec) Get(path string, opts ...api.PathOpts) *api.API {

	return s.addNewApi(path, "GET", opts)
}
func (s *spec) Post(path string, opts ...api.PathOpts) *api.API {

	return s.addNewApi(path, "POST", opts)
}
func (s *spec) Delete(path string, opts ...api.PathOpts) *api.API {

	return s.addNewApi(path, "DELETE", opts)
}
func (s *spec) Put(path string, opts ...api.PathOpts) *api.API {

	return s.addNewApi(path, "PUT", opts)
}
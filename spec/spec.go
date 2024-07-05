package spec

import (
	"log"
	"r5t/api"

	"github.com/getkin/kin-openapi/openapi3"
)

type spec struct {
	// some info, do not have apis
	root openapi3.T
	apis []*api.API
}

func (s *spec) GenerateDoc() {
	// s.root.Paths = new(openapi3.Paths)
	log.Println("一共有多少API:", len(s.apis))
	for _, v := range s.apis {
		s.root.AddOperation(v.Path, v.Method, v.PathItem.Options)
	}
}

type SpecOpts func(s *openapi3.Info)

// contact
func WithSpecContact(name string, mail string, url string) SpecOpts {
	return func(s *openapi3.Info) {
		s.Contact = &openapi3.Contact{
			Name:  name,
			Email: mail,
			URL:   url,
		}
	}
}

// title
func WithSpecTitle(title string) SpecOpts {
	return func(s *openapi3.Info) {
		s.Title = title
	}
}

// version
func WithSpecVersion(v string) SpecOpts {
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
		PathItem: &openapi3.PathItem{
			Options: &openapi3.Operation{},
		},
		Method: method,
		Path:   path,
	}
	s.apis = append(s.apis, newApi)
	newApi.DealPathItem(newApi.PathItem.Options, opts)
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

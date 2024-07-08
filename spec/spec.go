package spec

import (
	"r5t/api"
	"r5t/model"
	"r5t/path"

	"github.com/getkin/kin-openapi/openapi3"
)

type spec struct {
	// some info, do not have apis
	root openapi3.T
}

func (s *spec) MarshalJSON() ([]byte, error) {
	return s.root.MarshalJSON()
}

type SpecOpts func(s *openapi3.T)

// contact
func WithContact(name string, mail string, url string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Contact = &openapi3.Contact{
			Name:  name,
			Email: mail,
			URL:   url,
		}
	}
}

// title
func WithTitle(title string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Title = title
	}
}

// version
func WithVersion(v string) SpecOpts {
	return func(s *openapi3.T) {
		s.Info.Version = v
	}
}

// servers
func WithServer(url string, desc string) SpecOpts {
	return func(s *openapi3.T) {
		s.Servers = append(s.Servers, &openapi3.Server{
			URL:         url,
			Description: desc,
		})
	}
}

// 新建一个Spec
func NewSpec(specs ...SpecOpts) *spec {
	var s spec
	s.root = openapi3.T{
		Info: &openapi3.Info{},
		Components: &openapi3.Components{
			Schemas: make(openapi3.Schemas),
		},
	}
	s.root.Info = &openapi3.Info{}
	for _, v := range specs {
		v(&s.root)
	}
	s.root.OpenAPI = "3.0.0"
	return &s
}

func (s *spec) addNewApi(path string, method string, opts []path.PathOpts) *api.API {
	var newApi *api.API = &api.API{
		Operation: &openapi3.Operation{},
		Schemas:   &s.root.Components.Schemas,
	}
	s.root.AddOperation(path, method, newApi.Operation)
	newApi.DealPathItem(newApi.Operation, opts)
	return newApi
}

// some options function
func (s *spec) Get(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "GET", opts)
}
func (s *spec) Post(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "POST", opts)
}
func (s *spec) Delete(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "DELETE", opts)
}
func (s *spec) Put(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "PUT", opts)
}

// registerModel dev
func (s *spec) RegisterModel(m model.Model, opts ...model.ModelOpts) *spec {

	schema := new(openapi3.Schema)
	model.ParseModel(m.Type, schema)
	for _, v := range opts {
		v(schema)
	}
	s.root.Components.Schemas[m.Type.Name()] = &openapi3.SchemaRef{
		Value: schema,
	}
	return s
}

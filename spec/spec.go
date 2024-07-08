package spec

import (
	"r5t/api"
	"r5t/model"
	"r5t/path"
	"r5t/security"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	// some info, do not have apis
	root openapi3.T
}

func (s *Spec) MarshalJSON() ([]byte, error) {
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
func NewSpec(specs ...SpecOpts) *Spec {
	var s Spec
	s.root = openapi3.T{
		Info: &openapi3.Info{},
		Components: &openapi3.Components{
			Schemas:         make(openapi3.Schemas),
			SecuritySchemes: make(openapi3.SecuritySchemes),
		},
	}
	s.root.Info = &openapi3.Info{}
	for _, v := range specs {
		v(&s.root)
	}
	s.root.OpenAPI = "3.0.0"
	return &s
}

func (s *Spec) addNewApi(path string, method string, opts []path.PathOpts) *api.API {
	var newApi *api.API = &api.API{
		Operation: &openapi3.Operation{},
		Schemas:   &s.root.Components.Schemas,
	}
	s.root.AddOperation(path, method, newApi.Operation)
	newApi.DealPathItem(newApi.Operation, opts)
	return newApi
}

// some options function
func (s *Spec) Get(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "GET", opts)
}
func (s *Spec) Post(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "POST", opts)
}
func (s *Spec) Delete(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "DELETE", opts)
}
func (s *Spec) Put(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "PUT", opts)
}

// registerModel
func (s *Spec) RegisterModel(m model.Model, opts ...model.ModelOpts) *Spec {

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

// security
func (s *Spec) Security(opts ...security.SecurityModelOpts) *Spec {
	for _, v := range opts {
		ss := &openapi3.SecurityScheme{}
		name := v(ss)
		s.root.Components.SecuritySchemes[name] = &openapi3.SecuritySchemeRef{
			Value: ss,
		}
	}
	return s
}

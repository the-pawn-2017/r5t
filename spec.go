package r5t

import (
	"github.com/the-pawn-2017/r5t/api"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/path"
	"github.com/the-pawn-2017/r5t/security"
	"github.com/the-pawn-2017/r5t/spec"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	// some info, do not have apis
	root openapi3.T
}

func (s *Spec) MarshalJSON() ([]byte, error) {
	return s.root.MarshalJSON()
}
func (s *Spec) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(s.root)
}
func (s *Spec) UnMarshalYAML(src []byte) error {
	return yaml.Unmarshal(src, &s.root)
}

// 新建一个Spec
func NewSpec(specs ...spec.SpecOpts) *Spec {
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

// export
func (s *Spec) ExportData() *openapi3.T {
	return &s.root
}

func (s *Spec) addNewApi(path string, method string, opts []path.PathOpts) *api.API {
	var newApi *api.API = &api.API{
		Operation: &openapi3.Operation{
			Description: "",
		},
		Schemas: &s.root.Components.Schemas,
	}
	s.root.AddOperation(path, method, newApi.Operation)
	newApi.DealPathItem(newApi.Operation, opts)
	newApi.Operation.Responses = &openapi3.Responses{}
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
func (s *Spec) Options(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "OPTIONS", opts)
}
func (s *Spec) Patch(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "PATCH", opts)
}
func (s *Spec) Trace(path string, opts ...path.PathOpts) *api.API {

	return s.addNewApi(path, "TRACE", opts)
}
func (s *Spec) Head(path string, opts ...path.PathOpts) *api.API {
	return s.addNewApi(path, "HEAD", opts)
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

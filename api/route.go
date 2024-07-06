package api

import (
	"r5t/model"

	"github.com/getkin/kin-openapi/openapi3"
)

type API struct {
	// PathItem *openapi3.PathItem
	Operation *openapi3.Operation
}

type PathOpts func(s *openapi3.Operation)

func WithPathDesc(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc

	}
}

func WithPathSummary(desc string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Description = desc
	}
}

func WithPathTags(tags []string) PathOpts {
	return func(s *openapi3.Operation) {
		s.Tags = tags
	}
}

func WithSecurity() PathOpts {
	return func(s *openapi3.Operation) {
		// s.Description = desc
	}
}

// deal all
func (api *API) DealPathItem(operation *openapi3.Operation, opts []PathOpts) *API {

	for _, v := range opts {
		v(operation)
	}
	return api
}

/*
	 func (api *API) Add(path string, opts ...PathOpts) *API {
		api.Path = path
		api.Method = "put"
		api.pathItem.Options = &openapi3.Operation{}
		return api.dealPathItem(api.pathItem.Get, opts)
	}
*/
func (api *API) Request(m model.Model, opts ...model.ReqModelOpts) *API {
	api.Operation.RequestBody = &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: make(openapi3.Content),
		},
	}
	for _, v := range opts {
		v(api.Operation.RequestBody.Value)

	}
	content := api.Operation.RequestBody.Value.Content
	for k := range content {
		if k == model.ReqJSON {
			schema := new(openapi3.Schema)
			model.ParseModel(m.Type, schema)
			item := &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Value: nil,
				},
			}
			item.Schema.Value = schema
			// item.Schema.Value = &openapi3.Schema{
			// 	Description: "这个是描述，我明白",
			// 	Type:        &openapi3.Types{openapi3.TypeString},
			// 	Example:     "15",
			// }
			api.Operation.RequestBody.Value.Content[k] = item
		}
	}
	return api
}
func (api *API) Response(code int, m model.Model) *API {
	return api
}
func (api *API) Header(name string, value string) *API {
	return api
}

// registerModel dev
func (api *API) RegisterModel(model *model.Model, opts ...model.ModelOpts) {
	model.Options = opts
}

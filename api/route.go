package api

import (
	"log"
	"r5t/model"
	"r5t/req"
	"r5t/res"

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
func (api *API) Request(m model.Model, opts ...req.ReqModelOpts) *API {
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
		if k == req.ReqJSON {
			schema := new(openapi3.Schema)
			model.ParseModel(m.Type, schema)
			item := &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Value: nil,
				},
			}
			item.Schema.Value = schema
			api.Operation.RequestBody.Value.Content[k] = item
		}
	}
	return api
}
func (api *API) Response(code int, m model.Model, opts ...res.ResModelOpts) *API {
	resbody := &openapi3.Response{
		Content: make(openapi3.Content),
	}
	for _, v := range opts {
		v(resbody)
	}
	for k := range resbody.Content {
		if k == res.ReqJSON {
			log.Println("我进来了")
			schema := new(openapi3.Schema)
			model.ParseModel(m.Type, schema)
			item := &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Value: nil,
				},
			}
			item.Schema.Value = schema
			log.Println(item.Schema.Value)
			resbody.Content[k] = item
		}
	}

	api.Operation.AddResponse(code, resbody)
	return api
}
func (api *API) Header(name string, value string) *API {
	return api
}

// registerModel dev
func (api *API) RegisterModel(model *model.Model, opts ...model.ModelOpts) {
	model.Options = opts
}

package api

import (
	"log"
	"r5t/header"
	"r5t/model"
	"r5t/param"
	"r5t/req"
	"r5t/res"

	"github.com/getkin/kin-openapi/openapi3"
)

type API struct {
	// PathItem *openapi3.PathItem
	Operation *openapi3.Operation
}

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
// Deprecated: Use ReqJSON instead.
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
		if k == header.ApplicationJson {
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

// theRequest is json and form
func (api *API) ReqJSON(m model.Model, opts ...req.ReqModelOpts) *API {
	jsonContent := openapi3.NewContentWithJSONSchema(&openapi3.Schema{})
	api.Operation.RequestBody = &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: jsonContent,
		},
	}
	for _, v := range opts {
		v(api.Operation.RequestBody.Value)

	}
	model.ParseModel(m.Type, jsonContent[header.ApplicationJson].Schema.Value)
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

func (api *API) Param(opts ...param.ReqParamOpts) *API {
	api.Operation.Parameters = make(openapi3.Parameters, 0)
	for _, v := range opts {
		v(&api.Operation.Parameters)
	}
	return api
}

// registerModel dev
func (api *API) RegisterModel(model *model.Model, opts ...model.ModelOpts) {
	model.Options = opts
}

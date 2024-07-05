package api

import (
	"log"
	"r5t/model"

	"github.com/getkin/kin-openapi/openapi3"
)

type API struct {
	Path     string
	Param    string
	Method   string
	PathItem *openapi3.PathItem
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
	log.Println("api返回前:", api.PathItem.Options, api.PathItem.Description)
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
func (api *API) Request(code int, m model.Model) *API {
	return api
}
func (api *API) Response(code int, m model.Model) *API {
	return api
}
func (api *API) Header(name string, value string) *API {
	return api
}

// registerModel
func (api *API) RegisterModel(model *model.Model, opts ...model.ModelOpts) {
	model.Options = opts
}

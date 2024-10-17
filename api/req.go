package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/the-pawn-2017/r5t/header"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/req"
)

func (api *API) ReqCustom(m model.Model, header string, opts ...req.ReqModelOpts) *API {
	if m.Type == nil {
		return nil
	}
	jsonContent := openapi3.NewContentWithSchema(&openapi3.Schema{}, []string{header})
	api.Operation.RequestBody = &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: jsonContent,
		},
	}
	for _, v := range opts {
		v(api.Operation.RequestBody.Value)

	}
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, jsonContent[header].Schema.Value)
	} else {
		api.Operation.RequestBody.Ref = "#/components/schemas/" + m.Type.Name()
	}

	return api
}

func (api *API) ReqJSON(m model.Model, opts ...req.ReqModelOpts) *API {
	return api.ReqCustom(m, header.ApplicationJson, opts...)
}

func (api *API) ReqFormNoFile(m model.Model, opts ...req.ReqModelOpts) *API {
	return api.ReqCustom(m, header.ApplicationXWwwFormUrlencoded, opts...)
}

func (api *API) ReqFormWithFile(m model.Model, opts ...req.ReqModelOpts) *API {
	return api.ReqCustom(m, header.MultipartFormData, opts...)
}

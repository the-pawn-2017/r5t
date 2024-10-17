package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/the-pawn-2017/r5t/header"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/res"
)

func (api *API) ResCustom(code int, header string, m model.Model, opts ...res.ResModelOpts) *API {
	if m.Type == nil {
		return nil
	}
	resbody := openapi3.NewResponse()
	resbody.Content = openapi3.NewContent()
	resbody.Content[header] = openapi3.NewMediaType()
	resbody.Content[header].Schema = &openapi3.SchemaRef{
		Value: &openapi3.Schema{},
	}
	for _, v := range opts {
		v(resbody)
	}
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, resbody.Content[header].Schema.Value)
	} else {
		resbody.WithJSONSchemaRef(&openapi3.SchemaRef{
			Ref: "#/components/schemas/" + m.Type.Name(),
		})
	}
	api.Operation.AddResponse(code, resbody)
	return api
}

func (api *API) ResJSON(code int, m model.Model, opts ...res.ResModelOpts) *API {
	return api.ResCustom(code, header.ApplicationJson, m, opts...)
}

func (api *API) ResString(code int, opts ...res.ResModelOpts) *API {

	return api.ResCustom(code, header.TextPlain, model.ModelOf[string](), opts...)
}

package api

import (
	"github.com/the-pawn-2017/r5t/header"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/param"
	"github.com/the-pawn-2017/r5t/path"
	"github.com/the-pawn-2017/r5t/req"
	"github.com/the-pawn-2017/r5t/res"

	"github.com/getkin/kin-openapi/openapi3"
)

type API struct {
	// PathItem *openapi3.PathItem
	Operation *openapi3.Operation
	Schemas   *openapi3.Schemas
}

// deal all
func (api *API) DealPathItem(operation *openapi3.Operation, opts []path.PathOpts) *API {

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
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, jsonContent[header.ApplicationJson].Schema.Value)
	} else {
		api.Operation.RequestBody.Ref = "#/components/schemas/" + m.Type.Name()
	}

	return api
}

func (api *API) ReqCustom(m model.Model, header string, opts ...req.ReqModelOpts) *API {
	jsonContent := openapi3.NewContentWithJSONSchema(&openapi3.Schema{})
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

func (api *API) ReqFormNoFile(m model.Model, opts ...req.ReqModelOpts) *API {
	jsonContent := openapi3.NewContentWithSchema(&openapi3.Schema{}, []string{header.ApplicationXWwwFormUrlencoded})
	api.Operation.RequestBody = &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: jsonContent,
		},
	}
	for _, v := range opts {
		v(api.Operation.RequestBody.Value)

	}
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, jsonContent[header.ApplicationXWwwFormUrlencoded].Schema.Value)
	} else {
		api.Operation.RequestBody.Ref = "#/components/schemas/" + m.Type.Name()
	}

	return api
}

func (api *API) ReqFormWithFile(m model.Model, opts ...req.ReqModelOpts) *API {
	formContent := openapi3.NewContentWithFormDataSchema(&openapi3.Schema{})
	api.Operation.RequestBody = &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: formContent,
		},
	}
	for _, v := range opts {
		v(api.Operation.RequestBody.Value)

	}
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, formContent[header.MultipartFormData].Schema.Value)
	} else {
		api.Operation.RequestBody.Ref = "#/components/schemas/" + m.Type.Name()
	}

	return api
}

func (api *API) ResJSON(code int, m model.Model, opts ...res.ResModelOpts) *API {
	resbody := &openapi3.Response{
		Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
			Description: "",
		}),
	}
	for _, v := range opts {
		v(resbody)
	}
	if (*api.Schemas)[m.Type.Name()] == nil {
		model.ParseModel(m.Type, resbody.Content[header.ApplicationJson].Schema.Value)
	} else {
		resbody.WithJSONSchemaRef(&openapi3.SchemaRef{
			Ref: "#/components/schemas/" + m.Type.Name(),
		})
		//resbody.Content[header.ApplicationJson].Schema.Ref = "#/components/schemas/" + m.Type.Name()
	}
	api.Operation.AddResponse(code, resbody)
	return api
}

func (api *API) ResCustom(code int, header string, m model.Model, opts ...res.ResModelOpts) *API {
	resbody := &openapi3.Response{
		Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
			Description: "",
		}),
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
		//resbody.Content[header.ApplicationJson].Schema.Ref = "#/components/schemas/" + m.Type.Name()
	}
	api.Operation.AddResponse(code, resbody)
	return api
}

func (api *API) ResString(code int, opts ...res.ResModelOpts) *API {
	resbody := &openapi3.Response{
		Content: openapi3.NewContent(),
	}
	resbody.Content[header.TextPlain] = openapi3.NewMediaType()
	resbody.Content[header.TextPlain].Example = "there are some msg"
	var desc string
	resbody.Description = &desc
	for _, v := range opts {
		v(resbody)
	}

	api.Operation.AddResponse(code, resbody)
	return api
}

func (api *API) dealParam(name string, in string, opts []param.ReqParamOpts) *API {
	if len(api.Operation.Parameters) == 0 {
		api.Operation.Parameters = make(openapi3.Parameters, 0)
	}
	pList := &api.Operation.Parameters
	p := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:   in,
			Name: name,
			//Required: true,
		}}
	*pList = append(*pList, &p)

	for _, v := range opts {
		v(p.Value)
	}
	return api
}

func (api *API) NeedSecurify(tokenName string, require []string) *API {
	if api.Operation.Security == nil {
		api.Operation.Security = new(openapi3.SecurityRequirements)
	}
	*api.Operation.Security = append(*api.Operation.Security, openapi3.SecurityRequirement{
		tokenName: require,
	})
	return api
}

func (api *API) Path(name string, opts ...param.ReqParamOpts) *API {
	return api.dealParam(name, param.InPath, opts)
}

func (api *API) Cookie(name string, opts ...param.ReqParamOpts) *API {
	return api.dealParam(name, param.InCookie, opts)
}
func (api *API) Header(name string, opts ...param.ReqParamOpts) *API {
	return api.dealParam(name, param.InHeader, opts)
}
func (api *API) Query(name string, opts ...param.ReqParamOpts) *API {
	return api.dealParam(name, param.InQuery, opts)
}

// page
func (api *API) PageInQuery(pageName string, defaultPageNum int, pageSizeName string, defaultPageSizeNum int) *API {
	api.dealParam(pageName, param.InQuery, []param.ReqParamOpts{param.Desc("page index"), param.Default(defaultPageNum)})
	api.dealParam(pageSizeName, param.InQuery, []param.ReqParamOpts{param.Desc("page size"), param.Default(defaultPageSizeNum)})
	return api
}

/*
If my API does not meet your requirements, you can modify certain things yourself using append operations.
I hope it useful.
api has Operation and Schemas, it can be directly update.
like tests/TestAppend func
*/
func (api *API) Append(f func(api *API)) {
	f(api)
}

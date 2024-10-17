package api

import (
	"github.com/the-pawn-2017/r5t/param"

	"github.com/getkin/kin-openapi/openapi3"
)

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

// PageInQuery is a convenient function for adding page and pageSize parameters to the API.
func (api *API) PageInQuery(pageName string, defaultPageNum int, pageSizeName string, defaultPageSizeNum int) *API {
	api.dealParam(pageName, param.InQuery, []param.ReqParamOpts{param.Desc("page index"), param.Default(defaultPageNum)})
	api.dealParam(pageSizeName, param.InQuery, []param.ReqParamOpts{param.Desc("page size"), param.Default(defaultPageSizeNum)})
	return api
}

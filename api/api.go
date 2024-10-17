package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/the-pawn-2017/r5t/path"
)

type API struct {
	Operation *openapi3.Operation
	Schemas   *openapi3.Schemas
}

// deal all rest request
// 处理当前路径配置
func (api *API) DealPathItem(operation *openapi3.Operation, opts []path.PathOpts) *API {

	for _, v := range opts {
		v(operation)
	}
	return api
}

/*
If r5t API does not meet your requirements, you can modify certain things yourself using append operations.
I hope it useful.
api has Operation and Schemas, it can be directly update.
like tests/TestAppend func
*/
func (api *API) Append(f func(api *API)) {
	f(api)
}

package api

import "github.com/getkin/kin-openapi/openapi3"

func (api *API) NeedSecurify(tokenName string, require []string) *API {
	if api.Operation.Security == nil {
		api.Operation.Security = new(openapi3.SecurityRequirements)
	}
	*api.Operation.Security = append(*api.Operation.Security, openapi3.SecurityRequirement{
		tokenName: require,
	})
	return api
}

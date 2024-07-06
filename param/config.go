package param

import "github.com/getkin/kin-openapi/openapi3"

type ReqParamOpts func(*openapi3.Parameters)

// "query", "header", "path" or "cookie"
const (
	InQuery  = "query"
	InHeader = "header"
	InPath   = "path"
	InCookie = "cookie"
)

func WithQuery(name string, examples ...any) ReqParamOpts {
	return func(p *openapi3.Parameters) {
		*p = append(*p, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:   InQuery,
				Name: name,
			},
		})
	}
}
func WithPath(name string) ReqParamOpts {
	return func(p *openapi3.Parameters) {
		*p = append(*p, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:   InPath,
				Name: name,
			},
		})
	}
}
func WithHeader(name string) ReqParamOpts {
	return func(p *openapi3.Parameters) {
		*p = append(*p, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:   InHeader,
				Name: name,
			},
		})
	}
}
func WithCookie(name string) ReqParamOpts {
	return func(p *openapi3.Parameters) {
		*p = append(*p, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:   InCookie,
				Name: name,
			},
		})
	}
}

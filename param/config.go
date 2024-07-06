package param

import "github.com/getkin/kin-openapi/openapi3"

type ReqParamOpts func(*openapi3.Parameter)

// "query", "header", "path" or "cookie"
const (
	InQuery  = "query"
	InHeader = "header"
	InPath   = "path"
	InCookie = "cookie"
)

func WithExample(value string) ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Example = value
	}
}

func WithDesc(value string) ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Description = value
	}
}

func WithRequired() ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Required = true
	}
}

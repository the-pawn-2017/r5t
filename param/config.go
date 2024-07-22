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

func Example[T string | int | float64](value T) ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Example = value
	}
}

func Desc(value string) ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Description = value
	}
}

func Required() ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Required = true
	}
}
func Default[T string | bool | int | float64](v T) ReqParamOpts {
	return func(p *openapi3.Parameter) {
		p.Schema = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Default: nil,
			},
		}
		p.Schema.Value.Default = v

	}
}

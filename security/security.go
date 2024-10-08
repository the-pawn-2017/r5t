package security

import "github.com/getkin/kin-openapi/openapi3"

type SecurityModelOpts func(*openapi3.SecurityScheme) string

func HttpBasic(tokenName string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "http"
		ss.Scheme = "basic"
		return tokenName
	}
}

func ApiKey(tokenName string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "apiKey"
		ss.Scheme = "api_key"
		ss.In = "header"
		return tokenName
	}
}

func JWT(tokenName string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "apiKey"
		ss.Scheme = "bearer"
		ss.BearerFormat = "JWT"
		return tokenName
	}
}

const ImplicitFlow = "implicit"
const PasswordFlow = "password"
const ClientCredentialsFlow = "clientCredentials"
const AuthorizationCodeFlow = "authorizationCode"

func OAuth2Implicit(tokenName string, authorizationUrl string, opts ...scopesOpts) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "oauth2"
		ss.Flows = &openapi3.OAuthFlows{
			Implicit: &openapi3.OAuthFlow{
				AuthorizationURL: authorizationUrl,
				Scopes:           make(map[string]string),
			},
		}
		for _, v := range opts {
			v(&ss.Flows.AuthorizationCode.Scopes)
		}
		return tokenName
	}
}

type scopesOpts func(*map[string]string)

func AddScope(name string, desc string) scopesOpts {
	return func(scopes *map[string]string) {
		(*scopes)[name] = desc
	}
}

/*
	scopes: map[value]some info
*/
func OAuth2Code(tokenName string, authorizationUrl string, token string, opts ...scopesOpts) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "oauth2"
		ss.Flows = &openapi3.OAuthFlows{
			AuthorizationCode: &openapi3.OAuthFlow{
				AuthorizationURL: authorizationUrl,
				TokenURL:         token,
				Scopes:           make(map[string]string),
			},
		}
		for _, v := range opts {
			v(&ss.Flows.AuthorizationCode.Scopes)
		}
		return tokenName
	}
}

package security

import "github.com/getkin/kin-openapi/openapi3"

type SecurityModelOpts func(*openapi3.SecurityScheme) string

func WithHttpBasic(tokenName string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "http"
		ss.Scheme = "basic"
		return tokenName
	}
}

func WithApiKey(tokenName string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "apiKey"
		ss.Scheme = "api_key"
		ss.In = "header"
		return tokenName
	}
}

func WithJWT(tokenName string) SecurityModelOpts {
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

func WithOAuth2Implicit(tokenName string, authorizationUrl string, scopes map[string]string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "oauth2"
		ss.Flows = &openapi3.OAuthFlows{
			Implicit: &openapi3.OAuthFlow{
				AuthorizationURL: authorizationUrl,
				Scopes:           scopes,
			},
		}
		return tokenName
	}
}
func WithOAuth2Code(tokenName string, authorizationUrl string, token string, scopes map[string]string) SecurityModelOpts {
	return func(ss *openapi3.SecurityScheme) string {
		ss.Type = "oauth2"
		ss.Flows = &openapi3.OAuthFlows{
			AuthorizationCode: &openapi3.OAuthFlow{
				AuthorizationURL: authorizationUrl,
				TokenURL:         token,
				Scopes:           scopes,
			},
		}
		return tokenName
	}
}

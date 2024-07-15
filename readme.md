# r5t

> inspired by [a-h/rest](https://github.com/a-h/rest), testing code using `a-h/rest`

**DON'T TRY TO USE IT IN YOUR PROJECT, BECAUSE IT IN DEVELOPING**

## why create this repo?
1. I would like to implement more other features, such as support for GIN and ECHO.
2. Since many of my projects after that require REST API documentation, I'm more motivated to maintain it.
## version
v0.0.3
## todo
- [x] all components support and limit
- [x] param config, but no example and limit
- [x] Registering res&req model,now,it can use json,form, others type in developing.
- [x] Supporting OAuth2 , only code and implicit
- [x] register model
- [ ] complete test
- [ ] gin &echo support, now, echo can use r5t by some function, it's in [`example/echo`](./example/echo/echo.md)

... need more

## example:
go `/test/spec_test` view some example
```golang
type Test struct {
	A string
	B string `validate:"required"`
}
```
```golang
	s := spec.NewSpec()
	s.Security(
		security.WithOAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token",
			security.AddScope("openid", "OPENID IS USING FOR ID")),
	).
		Post("/gkd").NeedSecurify("ziteal", []string{"openid"}).
		ReqJSON(model.ModelOf[Test](), req.WithExample(Test{A: "A", B: "B"})).
		ResJSON(http.StatusOK, model.ModelOf[Test](), res.WithExample(Test{A: "A", B: "B"}))
```
## tools 
[swagger-ui-edit](https://editor-next.swagger.io/)
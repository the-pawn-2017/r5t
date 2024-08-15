# r5t

> It is a new implementation based on `go-openapi`。

**CAREFULLY USE IT IN YOUR PROJECT, BECAUSE IT IN DEVELOPING**

**I am currently testing with my own projects to refine R5T, expecting it to stabilize by the end of August. At that point, I will mark R5T as ready for official use, making it convenient for everyone.**



## why create this repo?
1. I would like to implement more other features, such as support for GIN and ECHO.
2. Since many of my projects after that require REST API documentation, I'm more motivated to maintain it.
## version
v0.2
## todo
- [x] all components support and limit
- [x] param config, but no example and limit
- [x] Registering res&req model,now,it can use json,form, others type in developing.
- [x] Supporting OAuth2 , only code and implicit
- [x] register model
- [ ] complete test
- [ ] full document for this repo
- [x] Support other web server,now,echo can use r5t by some function, it's in [`example/echo`](./example/echo/echo.md)

## some useful feature
### 1. fast pagination
```golang
s := r5t.NewSpec(spec.Title("pagination.yaml"))
s.Get("/test-pagination").PageInQuery("page", 1, "pageSize", 10).ResString(http.StatusOK, res.Example("hi"))
```
### 2. easy to use for OAuth2

### 3. concise and powerful API, like `Reqjson`,`ResJson`,`ResString`.
```golang
s := r5t.NewSpec(spec.Title("example reqString"))
s.Get("/test-resString").ResString(http.StatusOK, res.Example("hi!"))
```
## example:
go [`/test/spec_test`](/tests/spec_test.go) view some example
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
	)
	// than, you can use OAuth2 code mode now
	s.Post("/gkd").NeedSecurify("ziteal", []string{"openid"}).
		ReqJSON(model.ModelOf[Test](), req.WithExample(Test{A: "A", B: "B"})).
		ResJSON(http.StatusOK, model.ModelOf[Test](), res.WithExample(Test{A: "A", B: "B"}))
```
## tools 
[swagger-ui-edit](https://editor-next.swagger.io/)
> inspired by [a-h/rest](https://github.com/a-h/rest)
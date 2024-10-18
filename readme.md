# r5t

> 省流版：这是一个用go 代码生成swagger/openApi文档的库，不需要写注释！;这个库基于`go-openapi`，可以嵌入到web项目中。[文档连接](https://juejin.cn/user/272334613918430/posts)
> Automatically generated OpenAPI 3/swagger documentation via Go code, without relying on comments. It can be embedded in your web project.

install: `go get -u github.com/the-pawn-2017/r5t`

## why create this repo?
1. I would like to implement more other features, such as support for GIN and ECHO.
2. Since many of my projects after that require REST API documentation, I'm more motivated to maintain it.
## version
v0.5
## todo
- ✅ all openAPI/swagger components support and limit
- ✅ Registering res&req model,now,it can use json,form.
- ✅ Supporting OAuth2 , only code and implicit
- ✅ register model
- 🚧 complete test
- 🚧 full document for this repo
- ✅ Support other web server,now,echo can use `r5t` by some function, it's in [`example/echo`](./example/echo/echo.md)

## some useful feature
### 1. fast pagination
```golang
s := r5t.NewSpec(spec.Title("pagination.yaml"))
s.Get("/test-pagination").PageInQuery("page", 1, "pageSize", 10).ResString(http.StatusOK, res.Example("hi"))
```
### 2. easy to use for OAuth2
```golang
s := spec.NewSpec()
s.Security(
	security.WithOAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token",
	security.AddScope("openid", "OPENID IS USING FOR ID")),
)
```
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
### embed swagger-ui

```golang
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/the-pawn-2017/r5t"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/req"
	"github.com/the-pawn-2017/r5t/res"
	"github.com/the-pawn-2017/r5t/security"
	"github.com/the-pawn-2017/r5t/swaggerui"
)

type TestBasic struct {
	A string
	B string `validate:"required"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	s := r5t.NewSpec()
	s.Security(
		security.OAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token",
			security.AddScope("openid", "OPENID IS USING FOR ID")),
	).
		Post("/gkd").NeedSecurify("ziteal", []string{"openid"}).
		ReqJSON(model.ModelOf[TestBasic](), req.Example(TestBasic{A: "A", B: "B"})).
		ResJSON(http.StatusOK, model.ModelOf[TestBasic](), res.Example(TestBasic{A: "A", B: "B"}))
	e.GET("/swagger-test.json", func(c echo.Context) error {
		re, err := swaggerui.GenSpec(s)
		if err == nil {
			return c.JSONBlob(http.StatusOK, re)
		} else {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	})
	e.GET("/swagger/*", func(c echo.Context) error {
		paramStr := c.Param("*")
		kind, content, err := swaggerui.GetSwaggerUIFile("/swagger-test.json", paramStr)
		if err == nil {
			return c.Blob(http.StatusOK, kind, content)
		}
		return c.String(http.StatusInternalServerError, err.Error())
	})
	e.Start(":2333")
}
```


[`example/echo`](./example/echo/echo.md)

**CAREFULLY USE IT IN YOUR PROJECT, BECAUSE IT IN DEVELOPING**

**I am currently testing with my own projects to refine R5T, expecting it to stabilize by the end of August. At that point, I will mark R5T as ready for official use, making it convenient for everyone.**
## tools 
[swagger-ui-edit](https://editor-next.swagger.io/)
> inspired by [a-h/rest](https://github.com/a-h/rest)
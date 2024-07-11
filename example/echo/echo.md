# go-echo example
## if you want to use go-echo, there are example like this:
```golang
package main

import (
	"net/http"
	"r5t/model"
	"r5t/req"
	"r5t/res"
	"r5t/security"
	"r5t/spec"
	"r5t/swaggerui"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TestBasic struct {
	A string
	B string `validate:"required"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	s := spec.NewSpec()
	s.Security(
		security.WithOAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token",
			security.AddScope("openid", "OPENID IS USING FOR ID")),
	).
		Post("/gkd").NeedSecurify("ziteal", []string{"openid"}).
		ReqJSON(model.ModelOf[TestBasic](), req.WithExample(TestBasic{A: "A", B: "B"})).
		ResJSON(http.StatusOK, model.ModelOf[TestBasic](), res.WithExample(TestBasic{A: "A", B: "B"}))
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
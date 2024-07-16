# go-echo example
## if you want to use go-echo, there are example like this:
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
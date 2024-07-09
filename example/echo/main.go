package main

import (
	"net/http"
	"r5t/model"
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
	s.Security(security.WithOAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token", map[string]string{
		"openid": "OPENID",
		"3":      "4",
	})).Get("gkd").ResJSON(http.StatusOK, model.ModelOf[TestBasic]()).ResJSON(http.StatusOK, model.ModelOf[TestBasic](), res.WithExample(TestBasic{A: "A", B: "B"}))
	e.GET("/swagger-test.json", swaggerui.GenSpec(s))
	e.GET("/swagger/*", swaggerui.GenSwaggerUI("/swagger-test.json"))
	e.Start(":2333")
}

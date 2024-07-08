package test

import (
	"net/http"
	"r5t/model"
	"r5t/security"
	"r5t/spec"
	"testing"
)

func TestBasic(t *testing.T) {

}

func serveAuthTest(w http.ResponseWriter, r *http.Request) {
	type TestBasic struct {
		A string
		B string `validate:"required"`
	}
	s := spec.NewSpec()
	s.Security(security.WithOAuth2Code("ziteal", "http://10.45.8.189:8080/oauth/v2/authorize", "http://10.45.8.189:8080/oauth/v2/token", map[string]string{
		"openid": "OPENID",
		"3":      "4",
	})).Get("gkd").ResJSON(http.StatusOK, model.ModelOf[TestBasic]())
	b, _ := s.MarshalJSON()
	w.Write(b)
}

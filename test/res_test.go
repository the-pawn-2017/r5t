package test

import (
	"net/http"
	"r5t/api"
	"r5t/model"
	"r5t/res"
	"r5t/spec"
	"testing"
)

func TestResStruct(t *testing.T) {
	type TestModel struct {
		One   string
		Two   *string
		Three struct {
			A string
			B int64
		}
	}
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Get("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).
		Response(http.StatusOK, model.ModelOf[TestModel](), res.WithJSON(true, "一段说明"))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

package test

import (
	"net/http"
	"r5t/model"
	"r5t/path"
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
	s.Get("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"})).
		ResJSON(http.StatusOK, model.ModelOf[TestModel]())
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

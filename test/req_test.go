package test

import (
	"r5t/model"
	"r5t/path"
	"r5t/req"
	"r5t/spec"
	"testing"
)

func TestReqStruct(t *testing.T) {
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
		Request(model.ModelOf[TestModel](), req.WithJSON(true, "一段说明"))
	s.Post("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"}))
	s.Delete("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"}))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

func TestReqStructHasEmbed(t *testing.T) {
	type TestModel3 struct {
		C bool
	}
	type TestModel struct {
		One string
		TestModel3
	}
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Post("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"})).ReqJSON(model.ModelOf[TestModel](), req.WithRequired(true))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

func TestReqArray(t *testing.T) {
	type TestStringArr []string
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Post("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"})).ReqJSON(model.ModelOf[TestStringArr](), req.WithDesc("我真的佛啦！"))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

func TestReqJSON(t *testing.T) {
	type TestStringArr []string
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Post("test-gkd", path.WithDesc("A test api item, get function"), path.WithSummary("hi!"), path.WithTags([]string{"k1"})).ReqJSON(model.ModelOf[TestStringArr]())
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

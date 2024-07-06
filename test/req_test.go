package test

import (
	"r5t/api"
	"r5t/model"
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
	s.Get("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).
		Request(model.ModelOf[TestModel](), req.WithJSON(true, "一段说明"))
	s.Post("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	s.Delete("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
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
	s.Post("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).Request(model.ModelOf[TestModel](), req.WithJSON(true, "一段说明"))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

func TestReqArray(t *testing.T) {
	type TestStringArr []string
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Post("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).Request(model.ModelOf[TestStringArr](), req.WithJSON(true, "我真的佛啦！"))
	re, _ := s.MarshalJSON()
	t.Log(string(re))
}

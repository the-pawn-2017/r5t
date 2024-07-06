package test

import (
	"log"
	"r5t/param"
	"r5t/spec"
	"testing"
)

func Test_Query(t *testing.T) {
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Get("test-query").ParamQuery("gkd", param.WithDesc("我是一个描述"), param.WithExample("123"))
	re, _ := s.MarshalJSON()
	log.Println(string(re))
}

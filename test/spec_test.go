package test

import (
	"net/http"
	"r5t/api"
	"r5t/model"
	"r5t/req"
	"r5t/spec"
	"testing"
	"text/template"
)

func TestSpecGen(t *testing.T) {
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Get("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	re, _ := s.MarshalJSON()
	t.Log(re)
}

func TestSpecOfSwaggerUI(t *testing.T) {
	// serveHTML 是处理函数

	// create http server
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/test.json", serveJSON)
	http.ListenAndServe(":8000", nil)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	// 定义要传递给模板的数据

	// 加载 HTML 模板文件
	tmpl, err := template.ParseFiles("../swaggerui/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 使用数据渲染 HTML 模板并写入响应
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// serveHTML 是处理函数
func serveJSON(w http.ResponseWriter, r *http.Request) {
	type TestModel struct {
		One string
		Two string
	}
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Get("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).
		Request(model.ModelOf[TestModel](), req.WithJSON(true, "一段说明"))
	s.Post("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	s.Delete("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	re, _ := s.MarshalJSON()
	w.Write(re)
}

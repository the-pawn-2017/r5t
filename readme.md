# r5t

> inspired by [a-h/rest](https://github.com/a-h/rest), and some code using the repo.

**DON'T TRY TO USE IT IN YOUR PROJECT, BECAUSE IT IN DEVELOPING**

## why create this repo?
1. I would like to implement more other features, such as support for GIN and ECHO.
2. Since many of my projects after that require REST API documentation, I'm more motivated to maintain it.
## version
v0.0.2
## todo
- [ ] all components example and limit
- [x] param config, but no example and limit
- [x] res&req model,now,it can use json, others type in developing.
- [x] OAuth2 model, only code and implicit
- [x] register model
- [ ] complete test
- [ ] gin &echo support, now, echo can use r5t by some function, it's in `example/echo`
... need more

## example:
go `/test/spec_test` view some example
```golang
	func serveJSON(w http.ResponseWriter, r *http.Request) {
	type TestModel struct {
		One string
		Two string
	}
	s := spec.NewSpec(spec.WithTitle("test page"), spec.WithVersion("0.0.1"))
	s.Get("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"})).Request(model.ModelOf[TestModel](), model.WithReqJSON(true, "一段说明"))
	s.Post("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	s.Delete("test-gkd", api.WithPathDesc("A test api item, get function"), api.WithPathSummary("hi!"), api.WithPathTags([]string{"k1"}))
	re, _ := s.MarshalJSON()
	w.Write(re)
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
func TestSpecOfSwaggerUI(t *testing.T) {
	// serveHTML 是处理函数

	// create http server
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/test.json", serveJSON)
	http.ListenAndServe(":8000", nil)
}
```
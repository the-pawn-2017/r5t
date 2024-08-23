package tests

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/the-pawn-2017/r5t"
	"github.com/the-pawn-2017/r5t/api"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/param"
	"github.com/the-pawn-2017/r5t/req"
	"github.com/the-pawn-2017/r5t/res"
	"github.com/the-pawn-2017/r5t/spec"
)

func TestFormFile(t *testing.T) {
	type embedStruct struct {
		C string
	}
	type TestStruct struct {
		A           string
		B           string
		embedStruct `json:"-"`
	}
	s := r5t.NewSpec(spec.Version("3.0.0"), spec.Title("test"))
	s.Post("/test").ReqFormWithFile(model.ModelOf[TestStruct](), req.FormFile("GKD.txt", "I need a text file", true)).
		ResString(http.StatusOK, res.Desc("过"))
	err := s.ExportData().Validate(context.Background())
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

// still need to test, I will design some function to test the repo.
func TestAllMethods(t *testing.T) {
	// path := "all-methods.yaml"
	//b, _ := os.ReadFile("./specs" + path)
	type O struct {
		Ok  bool `json:"OK"`
		Ok2 bool `json:"-"`
	}

	s := r5t.NewSpec(spec.Title("all-methods.yaml")).RegisterModel(model.ModelOf[O]())
	s.Delete("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Get("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Head("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Options("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Patch("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Post("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Put("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Trace("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	log.Println(s.MarshalYAML())
	genDiff(s, "./specs/"+"001-all-methods.yaml", t)

}

func TestPath(t *testing.T) {
	s := r5t.NewSpec(spec.Title("params.yaml"))
	s.Get("/page").PageInQuery("page", 1, "pageSize", 10)
	s.Get("/param/{abc}").Path("abc", param.Default(1), param.Example(1))
	genDiff(s, "./specs/"+"002-path.yaml", t)
}

func TestAppend(t *testing.T) {
	s := r5t.NewSpec(spec.Title("params.yaml"))
	s.Get("/test-append").Append(func(api *api.API) {
		api.Operation.Tags = []string{"test_tag"}
		api.Operation.Responses = openapi3.NewResponses()
	})
	genDiff(s, "./specs/"+"004-append.yaml", t)
}

func TestResString(t *testing.T) {
	s := r5t.NewSpec(spec.Title("params.yaml"))
	s.Get("/test-resString").ResString(http.StatusOK, res.Example("hi!"))
	re, _ := s.MarshalYAML()
	log.Println(string(re))
	genDiff(s, "./specs/"+"005-resString.yaml", t)
}

func TestPagination(t *testing.T) {
	s := r5t.NewSpec(spec.Title("pagination.yaml"))
	s.Get("/test-pagination").PageInQuery("page", 1, "pageSize", 10).ResString(http.StatusOK, res.Example("hi"))
	genDiff(s, "./specs/"+"006-pagination.yaml", t)
}

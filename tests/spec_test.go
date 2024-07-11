package tests

import (
	"log"
	"net/http"
	"os"
	"r5t/model"
	"r5t/req"
	"r5t/spec"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/r3labs/diff/v3"
)

func TestExportYaml(t *testing.T) {
	type TestStruct struct {
		A string
		B string
	}
	s := spec.NewSpec()
	s.Post("/test").ReqFormWithFile(model.ModelOf[TestStruct](), req.WithFormFile("GKD.txt", "I need a text file", true))
	re, _ := s.MarshalYAML()
	log.Println(string(re))
}

func TestLoadYAML(t *testing.T) {
	content, _ := os.ReadFile("./specs/form-file.yaml")
	expectedSpec, err := openapi3.NewLoader().LoadFromData(content)
	log.Println(err, "load err")
	log.Println(expectedSpec.Paths, expectedSpec.Info)

}

// still need to test, I will design some function to test the repo.
func AllMethodsTest(t *testing.T) {
	// path := "all-methods.yaml"
	//b, _ := os.ReadFile("./specs" + path)
	type O struct {
		Ok bool `json:"OK"`
	}

	s := spec.NewSpec(spec.WithTitle("all-methods.yaml"), spec.WithVersion("0.0.0")).RegisterModel(model.ModelOf[O]())
	s.Connect("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Delete("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Get("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Head("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Options("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Patch("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Post("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Put("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Trace("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.MarshalJSON()
}

func TestFormFile(t *testing.T) {
	type TestStruct struct {
		A string
		B string
	}
	s := spec.NewSpec()
	s.Post("/test").ReqFormWithFile(model.ModelOf[TestStruct](), req.WithFormFile("GKD.txt", "I need a text file", true))
	content, _ := os.ReadFile("./specs/form-file.yaml")
	expectedSpec, _ := openapi3.NewLoader().LoadFromData(content)

	if re, err := diff.Diff(expectedSpec, s.ExportData()); len(re) > 0 && err == nil {
		t.Error(re)
	} else if err != nil {
		t.FailNow()
	}
}

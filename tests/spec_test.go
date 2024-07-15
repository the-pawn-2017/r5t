package tests

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/req"
	"github.com/the-pawn-2017/r5t/spec"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/r3labs/diff/v3"
)

// still need to test, I will design some function to test the repo.
func TestAllMethods(t *testing.T) {
	// path := "all-methods.yaml"
	//b, _ := os.ReadFile("./specs" + path)
	type O struct {
		Ok bool `json:"OK"`
	}

	s := spec.NewSpec(spec.WithTitle("all-methods.yaml"), spec.WithVersion("0.0.0")).RegisterModel(model.ModelOf[O]())
	s.Delete("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Get("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Head("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Options("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Patch("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Post("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Put("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	s.Trace("/connect").ResJSON(http.StatusOK, model.ModelOf[O]())
	genDiff(s, "./specs/"+"001-all-methods.yaml", t)

}

func TestFormFile(t *testing.T) {
	type TestStruct struct {
		A string
		B string
	}
	s := spec.NewSpec()
	s.Post("/test").ReqFormWithFile(model.ModelOf[TestStruct](), req.WithFormFile("GKD.txt", "I need a text file", true))
	genDiff(s, "./specs/"+"000-form-file.yaml", t)
}

func genDiff(spec1 *spec.Spec, fileName string, t *testing.T) {
	d, _ := os.Getwd()
	log.Println(d)
	content, fileErr := os.ReadFile(fileName)
	if fileErr != nil {
		// t.Log(fileErr)
		t.Fatal(fileErr)
		return
	}
	specFromFile, parseErr := openapi3.NewLoader().LoadFromData(content)
	if parseErr != nil {
		// t.Log(parseErr)
		t.Fatal(parseErr)
		return
	}

	if re, err := diff.Diff(specFromFile, spec1.ExportData()); len(re) > 0 && err == nil {
		// t.Log(re)
		t.Fatal(re)
		return
	} else if err != nil && len(re) == 0 {
		t.Fatal(err)
	}
}

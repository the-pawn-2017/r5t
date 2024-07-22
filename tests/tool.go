package tests

import (
	"log"
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/r3labs/diff/v3"
	"github.com/the-pawn-2017/r5t"
)

func genDiff(spec1 *r5t.Spec, fileName string, t *testing.T) {
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

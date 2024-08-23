package tests

import (
	"testing"

	jd "github.com/josephburnett/jd/lib"
	"github.com/the-pawn-2017/r5t"
)

func genDiff(spec1 *r5t.Spec, fileName string, t *testing.T) {
	if nowJSON, err := spec1.MarshalJSON(); err == nil {
		preJSONNode, _ := jd.ReadJsonFile(fileName)
		nowJSONNode, _ := jd.ReadJsonString(string(nowJSON))
		t.Log(nowJSONNode.Json())
		df := preJSONNode.Diff(nowJSONNode)
		if len(df) != 0 {
			t.Log(df)
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

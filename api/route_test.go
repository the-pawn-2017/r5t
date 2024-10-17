package api

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/req"
)

func TestReqJSON(t *testing.T) {
	api := &API{
		Operation: &openapi3.Operation{},
	}
	type TestJson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 测试正常情况
	m := model.ModelOf[TestJson]()
	opts := []req.ReqModelOpts{
		req.Desc("test config can be used"),
	}
	result := api.ReqJSON(m, opts...)
	assert.NotNil(t, result)
	assert.Equal(t,
		result.Operation,
		openapi3.Operation{})

	// 测试模型不存在的情况
	m2 := model.Model{}
	result2 := api.ReqJSON(m2, opts...)
	assert.Nil(t, result2)
}

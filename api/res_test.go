package api

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/the-pawn-2017/r5t/header"
	"github.com/the-pawn-2017/r5t/model"
	"github.com/the-pawn-2017/r5t/res"
)

func TestJson(t *testing.T) {
	api := &API{
		Operation: &openapi3.Operation{},
		Schemas:   &openapi3.Schemas{},
	}
	type TestJson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 测试正常情况
	m := model.ModelOf[TestJson]()
	opts := []res.ResModelOpts{
		res.Desc("test config can be used"),
	}
	result := api.ResJSON(200, m, opts...)
	schema := openapi3.SchemaRef{
		Value: &openapi3.Schema{},
	}
	model.ParseModel(m.Type, schema.Value)
	desc := "test config can be used"
	assert.Equal(t,
		*result.Operation.Responses.Value("200"),
		openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &desc,
				Content: openapi3.Content{
					header.ApplicationJson: &openapi3.MediaType{
						Schema: &schema,
					},
				},
			},
		})

	// 测试模型不存在的情况
	m2 := model.Model{}
	result2 := api.ResJSON(200, m2, opts...)
	assert.Nil(t, result2)
}

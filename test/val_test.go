package test

import (
	"log"
	"net/http"
	"r5t/model"
	"r5t/spec"
	"testing"
)

func TestVal(t *testing.T) {
	type A struct {
		X int `validate:"required,gte=0,lte=130"`
	}

	r := spec.NewSpec()
	r.RegisterModel(model.ModelOf[A]()).Get("gkd").ReqJSON(model.ModelOf[A]()).ResJSON(http.StatusOK, model.ModelOf[A]())
	re, _ := r.MarshalJSON()
	log.Println(string(re))
}

package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/daytwominus/iban-validation-web-go/iban"
)

var validator = iban.NewValidator(iban.IbalLenByCodeFileProvider{FilePath: "../iban/LenByCode.txt"})

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/iban/validate/#iban", ValidateIban),
		rest.Post("/iban/validate", ValidateMultiple),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func ValidateIban(w rest.ResponseWriter, req *rest.Request) {
	res := validator.Validate(req.PathParam("iban"))
	w.WriteJson(res)
}

type MultipleIbans struct {
	Ibans []string
}

func ValidateMultiple(w rest.ResponseWriter, r *rest.Request) {
	multiple := MultipleIbans{}
	err := r.DecodeJsonPayload(&multiple)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []iban.IbanValidationResult
	for _, s := range multiple.Ibans {
		res = append(res, validator.Validate(s))
	}

	w.WriteJson(res)
}

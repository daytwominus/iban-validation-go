package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/daytwominus/iban-validation-web-go/iban"
	"github.com/gorilla/mux"
)

const port = "8080"

// we use file as a source of iban lengths
var validator = iban.NewValidator(iban.IbanLenByCodeFileProvider{FilePath: "./iban/LenByCode.txt"})

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/iban/validate/{iban}", ValidateIban).Methods("GET")
	r.HandleFunc("/iban/validate", ValidateMultiple).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ValidateIban(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ibanString := params["iban"]
	res := validator.Validate(ibanString)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", res)
}

// MultipleIbans representing what can be passed in POST body
type MultipleIbans struct {
	Ibans []string
}

func ValidateMultiple(w http.ResponseWriter, r *http.Request) {
	multiple := MultipleIbans{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&multiple)

	if err != nil {
		log.Fatal(err) //todo: write wrong request
		return
	}

	var res []iban.IbanValidationResult
	for _, s := range multiple.Ibans {
		res = append(res, validator.Validate(s))
	}

	bytes, err := json.Marshal(res)

	fmt.Fprintf(w, "%v", string(bytes))
}

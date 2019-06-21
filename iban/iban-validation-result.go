package iban

import (
	"encoding/json"
	"fmt"
)

type IbanValidationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (r IbanValidationResult) String() string {
	data, _ := json.Marshal(r)
	return fmt.Sprintf("%s\n", data)
}

// IBAn validation functionality

package iban

import (
	"fmt"
	"math/big"
	"strings"
)

type Validator struct {
	lenByCodeProvider IbalLenByCodeProvider
}

// cretaes new validator. Specify lenByCodeProvider as a source of pairs (countryCode, IBAN-length)
func NewValidator(lenByCodeProvider IbalLenByCodeProvider) *Validator {
	return &Validator{
		lenByCodeProvider: lenByCodeProvider,
	}
}

// IBAN validation routine
func (validator *Validator) Validate(ibanString string) IbanValidationResult {
	ibanUp := strings.ToUpper(ibanString)

	if len(ibanUp) < 2 {
		return IbanValidationResult{Message: "IBAN string is too short"}
	}

	code := ibanUp[0:2]

	length := 0
	ok := false
	length, ok = validator.lenByCodeProvider.LenByCode()[code]

	if !ok {
		return IbanValidationResult{Message: fmt.Sprintf("No country code found: %v", code)}
	}

	if length != len(ibanUp) {
		return IbanValidationResult{Message: "IBAN string has incorrect length"}
	}

	ibanReversedString := fmt.Sprintf("%v%v", ibanUp[4:], ibanUp[0:4])
	ibanValueString := ""
	for _, char := range ibanReversedString {
		append := string(char)

		if char >= 'A' && char <= 'Z' {
			chValue := int(char)
			append = fmt.Sprintf("%v", chValue-55)
		}
		ibanValueString += string(append)
	}

	value, _ := new(big.Int).SetString(ibanValueString, 10)

	if value.Mod(value, big.NewInt(97)).Cmp(big.NewInt(1)) != 0 {
		return IbanValidationResult{Message: "Invalid 97 check remainder"}
	}

	return IbanValidationResult{Success: true}
}

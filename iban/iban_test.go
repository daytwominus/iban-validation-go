// testing iban validation

package iban

import "testing"

// for tests we use hardcoded
var validator = NewValidator(IbalLenByCodeHardcodedProvider{})

var validIbans = [...]string{
	"GB82WEST12345698765432",
	"AL35202111090000000001234567",
	"AD1400080001001234567890",
	"AT483200000012345864",
	"LB92000700000000123123456123",
	"MT31MALT01100000000000000000123",
	"MC5810096180790123456789085",
	"SC52BAHL01031234567890123456USD",
	"BY86AKBB10100000002966000000",
	"VG21PACG0000000123456789",
}

var invalidIbans = [...]string{
	"",
	"GB",
	"GB1",
	"GB82WEST1234569876543",
}

func TestValidIbans(t *testing.T) {
	for _, iban := range validIbans {
		res := validator.Validate(iban)
		if !res.Success {
			t.Errorf("iban %q has to be correct but validation failed", iban)
		}
	}
}

func TestInvalidIbans(t *testing.T) {
	for _, iban := range invalidIbans {
		res := validator.Validate(iban)
		if res.Success {
			t.Errorf("iban %q has to be incorrect but validation succeeded", iban)
		}
	}
}

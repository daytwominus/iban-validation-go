package main

import (
	"fmt"

	"github.com/daytwominus/iban-validation-web-go/iban"
)

func main() {
	//fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(iban.Validate("GB82WEST12345698765432"))
}

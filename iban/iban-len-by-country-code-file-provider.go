// reading IBAN lengths from file. Data is taken from nordea website:
// https://www.nordea.com/en/our-services/cashmanagement/iban-validator-and-information/iban-countries/

package iban

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// IbanLenByCodeFileProvider is responsible for
// getting pairs (country-code, iban-length) from file
type IbanLenByCodeFileProvider struct {
	FilePath string
	// lenByCoutryCodeFromFile is filled after file is read
	lenByCoutryCodeFromFile map[string]int
}

// LenByCode is an implementation of IbalLenByCodeProvider interface
func (p IbanLenByCodeFileProvider) LenByCode() map[string]int {
	if p.lenByCoutryCodeFromFile != nil {
		return p.lenByCoutryCodeFromFile
	}

	file, err := os.Open(p.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ret map[string]int
	ret = make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		array := regexp.MustCompile("[\\\t\\ ]+").Split(line, -1)
		num, _ := strconv.ParseInt(array[0], 10, 64)
		ret[array[1]] = int(num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	p.lenByCoutryCodeFromFile = ret
	return ret
}

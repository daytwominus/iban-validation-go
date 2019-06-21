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

type IbalLenByCodeFileProvider struct {
	FilePath string
}

var lenByCoutryCodeFromFile map[string]int

func (p IbalLenByCodeFileProvider) LenByCode() map[string]int {
	if lenByCoutryCodeFromFile != nil {
		return lenByCoutryCodeFromFile
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

	lenByCoutryCodeFromFile = ret
	return ret
}

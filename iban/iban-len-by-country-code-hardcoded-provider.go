// Hardcoded IBAN lengths depending on country code. The list is valid on 2019-06-19
package iban

// IbalLenByCodeHardcodedProvider is responsible for
// getting hardcoded pairs (country-code, iban-length)
type IbalLenByCodeHardcodedProvider struct {
}

// LenByCode is an implementation of IbalLenByCodeProvider interface
func (p IbalLenByCodeHardcodedProvider) LenByCode() map[string]int {
	return lenByCoutryCode
}

// hardcoded values
var lenByCoutryCode = map[string]int{
	"AL": 28,
	"DZ": 24,
	"AD": 24,
	"AO": 25,
	"AT": 20,
	"AZ": 28,
	"BH": 22,
	"BY": 28,
	"BE": 16,
	"BJ": 28,
	"BA": 20,
	"BR": 29,
	"VG": 24,
	"BG": 22,
	"BF": 27,
	"BI": 16,
	"CM": 27,
	"CV": 25,
	"FR": 27,
	"CG": 27,
	"CR": 21,
	"HR": 21,
	"CY": 28,
	"CZ": 24,
	"DK": 18,
	"DO": 28,
	"EG": 27,
	"EE": 20,
	"FO": 18,
	"FI": 18,
	"GA": 27,
	"GE": 22,
	"DE": 22,
	"GI": 23,
	"GR": 27,
	"GL": 18,
	"GT": 28,
	"GG": 22,
	"HU": 28,
	"IS": 26,
	"IR": 26,
	"IQ": 23,
	"IE": 22,
	"IM": 22,
	"IL": 23,
	"IT": 27,
	"CI": 28,
	"JE": 22,
	"JO": 30,
	"KZ": 20,
	"XK": 20,
	"KW": 30,
	"LV": 21,
	"LB": 28,
	"LI": 21,
	"LT": 20,
	"LU": 20,
	"MK": 19,
	"MG": 27,
	"ML": 28,
	"MT": 31,
	"MR": 27,
	"MU": 30,
	"MD": 24,
	"MC": 27,
	"ME": 22,
	"MZ": 25,
	"NL": 18,
	"NO": 15,
	"PK": 24,
	"PS": 29,
	"PL": 28,
	"PT": 25,
	"QA": 29,
	"RO": 24,
	"LC": 32,
	"SM": 27,
	"ST": 25,
	"SA": 24,
	"SN": 28,
	"RS": 22,
	"SC": 31,
	"SK": 24,
	"SI": 19,
	"ES": 24,
	"SE": 24,
	"CH": 21,
	"TL": 23,
	"TN": 24,
	"TR": 26,
	"UA": 29,
	"AE": 23,
	"GB": 22,
	"VA": 22,
}

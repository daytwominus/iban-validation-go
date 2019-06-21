package iban

// IbalLenByCodeProvider interface
// Defines getting the IBAN length based on country code.
// This interface is added just for case when multiple sources are available
// (http calls, file, hardcoded etc).
type IbalLenByCodeProvider interface {
	LenByCode() map[string]int
}

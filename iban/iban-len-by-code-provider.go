package iban

type IbalLenByCodeProvider interface {
	LenByCode() map[string]int
}

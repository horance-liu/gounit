package gounit

type unit interface {
	factor() Amount
	base() unit
}
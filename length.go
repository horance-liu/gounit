package gounit

type lengthUnit uint

const (
	inch lengthUnit = 1
	feet = 12 * inch
	yard = 3 * feet
	mile = 1760 * yard 
)

func (u lengthUnit)	factor() Amount {
	return Amount(u)
}

func (u lengthUnit) base() unit {
	return inch
}

func (u lengthUnit) baseStr() string {
	return "INCH"
}

type Length struct {
	quantity
}

func (l Length) Equal(r Length) bool {
	return l.quantity.equal(r.quantity)
}

func (l Length) NotEqual(r Length) bool {
	return l.quantity.notEqual(r.quantity)
}

func (l Length) Plus(r Length) Length {
	return Length{l.quantity.plus(r.quantity)}
}

func (l Length) Minus(r Length) (Length, bool) {
	if q, ok := l.quantity.minus(r.quantity); ok {
		return Length{q}, true
	}
	return Length{}, false
}

func Inch(amount Amount) Length {
	return Length{quantity{amount, inch}}
}

func Feet(amount Amount) Length {
	return Length{quantity{amount, feet}}
}

func Mile(amount Amount) Length {
	return Length{quantity{amount, mile}}
}

func Yard(amount Amount) Length {
	return Length{quantity{amount, yard}}
}
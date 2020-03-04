package quantity

type Amount uint
type LengthUnit uint

const (
	inch LengthUnit = 1
	feet = 12 * inch
	yard = 3 * feet
	mile = 1760 * yard 
)

type Length struct {
	amount Amount
	unit LengthUnit
}

func (l Length) amountInBaseUnit() Amount {
	return l.amount * Amount(l.unit)
}

func (l Length) Equal(r Length) bool {
	return l.amountInBaseUnit() == r.amountInBaseUnit()
}

func (l Length) NotEqual(r Length) bool {
	return !(l.Equal(r))
}

func Inch(amount Amount) Length {
	return Length{amount, inch}
}

func Feet(amount Amount) Length {
	return Length{amount, feet}
}

func Mile(amount Amount) Length {
	return Length{amount, mile}
}

func Yard(amount Amount) Length {
	return Length{amount, yard}
}
package quantity


type Amount uint
type LengthUnit uint

const (
	yard LengthUnit = 1
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

func Mile(amount Amount) Length {
	return Length{amount, mile}
}

func Yard(amount Amount) Length {
	return Length{amount, yard}
}
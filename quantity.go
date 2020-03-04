package gounit

type quantity struct {
	amount Amount
	unit unit
}

func (v quantity) amountInBaseUnit() Amount {
	return v.amount * Amount(v.unit.factor())
}

func (v quantity) equal(r quantity) bool {
	return v.amountInBaseUnit() == r.amountInBaseUnit()
}

func (v quantity) notEqual(r quantity) bool {
	return !(v.equal(r))
}

func (v quantity) plus(r quantity) quantity {
	return quantity{v.amountInBaseUnit() + r.amountInBaseUnit(), v.unit.base()}
}

func (v quantity) minus(r quantity) (quantity, bool) {
	if (v.amountInBaseUnit() < r.amountInBaseUnit()) {
		return quantity{}, false
	}
	return quantity{v.amountInBaseUnit() - r.amountInBaseUnit(), v.unit.base()}, true
}
package quantity

type VolumeUnit uint

const (
	tsp VolumeUnit = 1
	tbsp = 3 * tsp
	oz   = 2 * tbsp
)

type Volume struct {
	amount Amount
	unit VolumeUnit
}

func (v Volume) amountInBaseUnit() Amount {
	return v.amount * Amount(v.unit)
}

func (v Volume) Equal(r Volume) bool {
	return v.amountInBaseUnit() == r.amountInBaseUnit()
}

func (v Volume) NotEqual(r Volume) bool {
	return !(v.Equal(r))
}

func (v Volume) Plus(r Volume) Volume {
	return Volume{v.amountInBaseUnit() + r.amountInBaseUnit(), tsp}
}

func (v Volume) Minus(r Volume) (Volume, bool) {
	if (v.amountInBaseUnit() < r.amountInBaseUnit()) {
		return Volume{}, false
	}
	return Volume{v.amountInBaseUnit() - r.amountInBaseUnit(), tsp}, true
}

func Tsp(amount Amount) Volume {
	return Volume{amount, tsp}
}

func Tbsp(amount Amount) Volume {
	return Volume{amount, tbsp}
}

func Oz(amount Amount) Volume {
	return Volume{amount, oz}
}

package gounit

type volumeUnit uint

const (
	tsp volumeUnit = 1
	tbsp = 3 * tsp
	oz   = 2 * tbsp
)

func (u volumeUnit)	factor() Amount {
	return Amount(u)
}

func (u volumeUnit) base() unit {
	return tsp
}

func (u volumeUnit) baseStr() string {
	return "TSP"
}

type Volume struct {
	quantity
}

func (v Volume) Equal(r Volume) bool {
	return v.quantity.equal(r.quantity)
}

func (v Volume) NotEqual(r Volume) bool {
	return v.quantity.notEqual(r.quantity)
}

func (v Volume) Plus(r Volume) Volume {
	return Volume{v.quantity.plus(r.quantity)}
}

func (v Volume) Minus(r Volume) (Volume, bool) {
	if q, ok := v.quantity.minus(r.quantity); ok {
		return Volume{q}, true
	}
	return Volume{}, false
}

func Tsp(amount Amount) Volume {
	return Volume{quantity{amount, tsp}}
}

func Tbsp(amount Amount) Volume {
	return Volume{quantity{amount, tbsp}}
}

func Oz(amount Amount) Volume {
	return Volume{quantity{amount, oz}}
}
